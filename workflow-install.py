#!/usr/bin/env python3
# encoding: utf-8
#
# Copyright (c) 2013 <deanishe@deanishe.net>.
#
# MIT Licence. See http://opensource.org/licenses/MIT
#
# Created on 2013-11-01
#

# https://gist.github.com/cage1016/a633149f4672fb40ecb4e16e04664ff7

"""workflow-install [options] [<workflow-directory>...]
Install Alfred workflow(s).
You can specify where to install by specifying an Alfred version
with --alfred or a specific directory with -w or in ~/.workflow-install.json
By default, it should install in the latest version of Alfred.
If <workflow-directory> is not specified, the script will search the
current working directory recursively for a workflow (a directory
containing an `info.plist` file).
Usage:
    workflow-install [-v|-q|-d] [-a <version>] [-s] [-w <directory>]
        [<workflow-directory>...]
    workflow-install (-h|--help)
Options:
    -a, --alfred=<version>          version of Alfred to install workflow
                                    in, e.g. "3" or "4"
    -s, --symlink                   symlink workflow directory instead of
                                    copying it
    -w, --workflows=<directory>     where to install workflows
    -V, --version                   show version number and exit
    -h, --help                      show this message and exit
    -q, --quiet                     show error messages and above
    -v, --verbose                   show info messages and above
    -d, --debug                     show debug messages
"""



import sys
import os
import logging
import logging.handlers
import json
import plistlib
import shutil
import subprocess

__version__ = "0.4.0"
__author__ = "deanishe@deanishe.net"


log = None

DEFAULT_LOG_LEVEL = logging.WARNING
# LOGPATH = os.path.expanduser('~/Library/Logs/MyScripts.log')
# LOGSIZE = 1024 * 1024 * 5  # 5 megabytes

CONFIG_PATH = os.path.expanduser('~/.workflow-install.json')
DEFAULT_CONFIG = dict(workflows_directory='')

ALFRED_PREFS = os.path.expanduser(
    '~/Library/Application Support/Alfred/prefs.json')
ALFRED3_PREFS = os.path.expanduser(
    '~/Library/Preferences/com.runningwithcrayons.'
    'Alfred-Preferences-3.plist')

DEFAULT_DIR = os.path.expanduser('~/Library/Application Support/Alfred')
DEFAULT_DIR3 = os.path.expanduser('~/Library/Application Support/Alfred 3')


class TechnicolorFormatter(logging.Formatter):
    """
    Prepend level name to any message not level logging.INFO.
    Also, colour!
    """

    BLACK, RED, GREEN, YELLOW, BLUE, MAGENTA, CYAN, WHITE = list(range(8))

    RESET = "\033[0m"
    COLOUR_BASE = "\033[1;{:d}m"
    BOLD = "\033[1m"

    LEVEL_COLOURS = {
        logging.DEBUG: BLUE,
        logging.INFO: WHITE,
        logging.WARNING: YELLOW,
        logging.ERROR: MAGENTA,
        logging.CRITICAL: RED,
    }

    def __init__(self, fmt=None, datefmt=None, technicolor=True):
        logging.Formatter.__init__(self, fmt, datefmt)
        self.technicolor = technicolor
        self._isatty = sys.stderr.isatty()

    def format(self, record):
        if record.levelno == logging.INFO:
            msg = logging.Formatter.format(self, record)
            return msg
        if self.technicolor and self._isatty:
            colour = self.LEVEL_COLOURS[record.levelno]
            bold = (False, True)[record.levelno > logging.INFO]
            levelname = self.colourise('{:9s}'.format(record.levelname),
                                       colour, bold)
        else:
            levelname = '{:9s}'.format(record.levelname)
        return (levelname + logging.Formatter.format(self, record))

    def colourise(self, text, colour, bold=False):
        colour = self.COLOUR_BASE.format(colour + 30)
        output = []
        if bold:
            output.append(self.BOLD)
        output.append(colour)
        output.append(text)
        output.append(self.RESET)
        return ''.join(output)


# console output
console = logging.StreamHandler()
formatter = TechnicolorFormatter('%(message)s')
console.setFormatter(formatter)
console.setLevel(logging.DEBUG)

log = logging.getLogger('')
# log.addHandler(logfile)
log.addHandler(console)


def read_plist(path):
    """Convert plist to XML and read its contents."""
    cmd = [b'plutil', b'-convert', b'xml1', b'-o', b'-', path]
    xml = subprocess.check_output(cmd)
    return plistlib.readPlistFromString(xml)


def get_workflow_directory(version=None):
    """Return path to Alfred's workflow directory."""
    dirs = []
    if version == '3' or version is None:
        dirs.append(DEFAULT_DIR3)
        if os.path.exists(ALFRED3_PREFS):
            prefs = read_plist(ALFRED3_PREFS)
            dirs.append(prefs.get('syncfolder'))

    if version != '3':
        dirs.append(DEFAULT_DIR)
        if os.path.exists(ALFRED_PREFS):
            with open(ALFRED_PREFS, 'rb') as fp:
                prefs = json.load(fp)

            if not version:
                s = prefs.get('current')
                log.debug('workflow sync dir: %r', s)
                return os.path.join(s, 'workflows')

            dirs.append(prefs.get('syncfolders', {}).get(version, ''))

    for p in dirs[::-1]:
        if not p:
            continue

        p = os.path.expanduser(p)

        # Alfred preserves syncdir setting even if directory no longer exists.
        # In this case, Alfred falls back to using its default directory in
        # ~/Library/Application Support
        if os.path.exists(p):
            syncdir = p
            break
    else:
        log.debug('Alfred sync folder not found')
        return None

    wf_dir = os.path.join(syncdir, 'Alfred.alfredpreferences/workflows')
    log.debug('workflow sync dir : %r', wf_dir)

    if os.path.exists(wf_dir):
        log.debug('workflow directory retrieved from Alfred preferences')
        return wf_dir

    log.debug('Alfred.alfredpreferences/workflows not found')
    return None


def find_workflow_dir(dirpath):
    """Recursively search `dirpath` for a workflow.
    A workflow is a directory containing an `info.plist` file.
    """
    for root, _, filenames in os.walk(dirpath):
        if 'info.plist' in filenames:
            log.debug('Workflow found at %r', root)
            return root

    return None


def printable_path(dirpath):
    """Replace $HOME with ~."""
    return dirpath.replace(os.getenv('HOME'), '~')


def load_config():
    """Load configuration from file."""
    if not os.path.exists(CONFIG_PATH):
        with open(CONFIG_PATH, 'wb') as file:
            json.dump(DEFAULT_CONFIG, file)
            return DEFAULT_CONFIG

    with open(CONFIG_PATH) as file:
        return json.load(file)


def install_workflow(workflow_dir, install_base, symlink=False):
    """Install workflow at `workflow_dir` under directory `install_base`."""
    if symlink:
        log.debug("Linking workflow at %r to %r", workflow_dir, install_base)
    else:
        log.debug("Installing workflow at %r to %r",
                  workflow_dir, install_base)

    infopath = os.path.join(workflow_dir, 'info.plist')
    if not os.path.exists(infopath):
        log.error('info.plist not found : %s', infopath)
        return False

    with open(infopath,'rb') as fp:
        info=plistlib.load(fp)

    name = info['name']
    bundleid = info['bundleid']

    if not bundleid:
        log.error('Bundle ID is not set : %s', infopath)
        return False

    install_path = os.path.join(install_base, bundleid)

    action = ('Installing', 'Linking')[symlink]
    log.info('%s workflow `%s` to `%s` ...',
             action, name, printable_path(install_path))

    # delete existing workflow
    if os.path.exists(install_path) or os.path.lexists(install_path):

        log.info('Deleting existing workflow ...')

        if os.path.islink(install_path) or os.path.isfile(install_path):
            os.unlink(install_path)
        elif os.path.isdir(install_path):
            log.info('Directory : %s', install_path)
            shutil.rmtree(install_path)
        else:
            log.info('Something else : %s', install_path)
            os.unlink(install_path)

    # Symlink or copy workflow to destination
    if symlink:
        relpath = os.path.relpath(workflow_dir, os.path.dirname(install_path))
        log.debug('relative path : %r', relpath)
        os.symlink(relpath, install_path)
    else:
        shutil.copytree(workflow_dir, install_path)

    return True


def main(args=None):
    """Run program."""
    from docopt import docopt
    args = docopt(__doc__, version=__version__)

    if args.get('--verbose'):
        log.setLevel(logging.INFO)
    elif args.get('--quiet'):
        log.setLevel(logging.ERROR)
    elif args.get('--debug'):
        log.setLevel(logging.DEBUG)
    else:
        log.setLevel(DEFAULT_LOG_LEVEL)

    log.debug("Set log level to %s" %
              logging.getLevelName(log.level))

    log.debug('args : \n%s', args)

    workflows_directory = (
        args.get('--workflows')
        or get_workflow_directory(version=args.get('--alfred'))
        or load_config().get('workflows_directory')
    )

    if not workflows_directory:
        log.error("You didn't specify where to install the workflow(s).\n"
                  "Try -w workflow/install/path or -h for more info.")
        return 1
    workflows_directory = os.path.expanduser(workflows_directory)

    # Ensure workflows_directory is Unicode
    if not isinstance(workflows_directory, str):
        workflows_directory = str(workflows_directory, 'utf-8')

    workflow_paths = args.get('<workflow-directory>')

    if not workflow_paths:
        cwd = os.getcwd()
        wfdir = find_workflow_dir(cwd)
        if not wfdir:
            log.critical('No workflow found under %r', cwd)
            return 1
        workflow_paths = [wfdir]
    errors = False

    for path in workflow_paths:
        if not isinstance(path, str):
            path = str(path, 'utf-8')
        path = os.path.abspath(path)
        if not os.path.exists(path):
            log.error('Directory does not exist : %s', path)
            continue
        if not os.path.isdir(path):
            log.error('Not a directory : %s', path)
            continue
        if not install_workflow(path, workflows_directory,
                                args.get('--symlink')):
            errors = True

    if errors:
        return 1
    return 0


if __name__ == '__main__':
    sys.exit(main(sys.argv[1:]))