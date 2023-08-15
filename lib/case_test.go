package lib

import (
	"testing"
)

func TestCase_Camel(t *testing.T) {
	type fields struct {
		ChangeCaser
	}
	type args struct {
		m map[string]string
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
	}{
		{
			name: "camel",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Camel
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "howAreYou",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "thisStringShouldBeInCamelCase",
					"I got intern at geeksforgeeks":               "iGotInternAtGeeksforgeeks",
					"Here comes the garden":                       "hereComesTheGarden",
					"The quick brown fox jumps over the lazy dog": "theQuickBrownFoxJumpsOverTheLazyDog",
				},
			},
		},
		{
			name: "pascal",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Pascal
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "HowAreYou",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "ThisStringShouldBeInCamelCase",
					"I got intern at geeksforgeeks":               "IGotInternAtGeeksforgeeks",
					"Here comes the garden":                       "HereComesTheGarden",
					"The quick brown fox jumps over the lazy dog": "TheQuickBrownFoxJumpsOverTheLazyDog",
				},
			},
		},
		{
			name: "constant",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Constant
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "HOW_A_RE_YO_U",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "THIS_STRING_SHOULD_BE_IN_CAMEL_CASE",
					"I got intern at geeksforgeeks":               "I_GOT_INTERN_AT_GEEKSFORGEEKS",
					"Here comes the garden":                       "HERE_COMES_THE_GARDEN",
					"The quick brown fox jumps over the lazy dog": "THE_QUICK_BROWN_FOX_JUMPS_OVER_THE_LAZY_DOG",
				},
			},
		},
		{
			name: "dot",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Dot
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "how.a.re.yo.u",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "this.string.should.be.in.camel.case",
					"I got intern at geeksforgeeks":               "i.got.intern.at.geeksforgeeks",
					"Here comes the garden":                       "here.comes.the.garden",
					"The quick brown fox jumps over the lazy dog": "the.quick.brown.fox.jumps.over.the.lazy.dog",
				},
			},
		},
		{
			name: "lower",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Lower
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "how are you",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "this string should be in camel case",
					"I got intern at geeksforgeeks":               "i got intern at geeksforgeeks",
					"Here comes the garden":                       "here comes the garden",
					"The quick brown fox jumps over the lazy dog": "the quick brown fox jumps over the lazy dog",
				},
			},
		},
		{
			name: "lcfirst",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Lcfirst
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "how aRe yoU",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "tHIS STRING SHOULD BE IN CAMEL CASE",
					"I got intern at geeksforgeeks":               "i got intern at geeksforgeeks",
					"Here comes the garden":                       "here comes the garden",
					"The quick brown fox jumps over the lazy dog": "the quick brown fox jumps over the lazy dog",
				},
			},
		},
		{
			name: "no",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().No
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "how a re yo u",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "this string should be in camel case",
					"I got intern at geeksforgeeks":               "i got intern at geeksforgeeks",
					"Here comes the garden":                       "here comes the garden",
					"The quick brown fox jumps over the lazy dog": "the quick brown fox jumps over the lazy dog",
					"this-is-a-test":                              "this is a test",
				},
			},
		},
		{
			name: "param",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Param
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "how-a-re-yo-u",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "this-string-should-be-in-camel-case",
					"I got intern at geeksforgeeks":               "i-got-intern-at-geeksforgeeks",
					"Here comes the garden":                       "here-comes-the-garden",
					"The quick brown fox jumps over the lazy dog": "the-quick-brown-fox-jumps-over-the-lazy-dog",
				},
			},
		},
		{
			name: "path",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Path
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "how/a/re/yo/u",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "this/string/should/be/in/camel/case",
					"I got intern at geeksforgeeks":               "i/got/intern/at/geeksforgeeks",
					"Here comes the garden":                       "here/comes/the/garden",
					"The quick brown fox jumps over the lazy dog": "the/quick/brown/fox/jumps/over/the/lazy/dog",
				},
			},
		},
		{
			name: "sentence",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Sentence
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":               "How are you",
					"HELLO world! HOW are YoU?": "Hello world! how are you?",
					"here's an Example":         "Here's an example",
					"to change the case of your text, paste it here and press the corresponding buttons below": "To change the case of your text, paste it here and press the corresponding buttons below",
				},
			},
		},
		{
			name: "snake",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Snake
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "how_a_re_yo_u",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "this_string_should_be_in_camel_case",
					"I got intern at geeksforgeeks":               "i_got_intern_at_geeksforgeeks",
					"Here comes the garden":                       "here_comes_the_garden",
					"The quick brown fox jumps over the lazy dog": "the_quick_brown_fox_jumps_over_the_lazy_dog",
				},
			},
		},
		{
			name: "swap",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Swap
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "HOW ArE YOu",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "this string should be in camel case",
					"I got intern at geeksforgeeks":               "i GOT INTERN AT GEEKSFORGEEKS",
					"Here comes the garden":                       "hERE COMES THE GARDEN",
					"The quick brown fox jumps over the lazy dog": "tHE QUICK BROWN FOX JUMPS OVER THE LAZY DOG",
				},
			},
		},
		{
			name: "title",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Title
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "How A Re Yo U",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "This String Should Be In Camel Case",
					"I got intern at geeksforgeeks":               "I Got Intern At Geeksforgeeks",
					"Here comes the garden":                       "Here Comes The Garden",
					"The quick brown fox jumps over the lazy dog": "The Quick Brown Fox Jumps Over The Lazy Dog",
				},
			},
		},
		{
			name: "upper",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Upper
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "HOW ARE YOU",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "THIS STRING SHOULD BE IN CAMEL CASE",
					"I got intern at geeksforgeeks":               "I GOT INTERN AT GEEKSFORGEEKS",
					"Here comes the garden":                       "HERE COMES THE GARDEN",
					"The quick brown fox jumps over the lazy dog": "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG",
				},
			},
		},
		{
			name: "ucfirst",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Ucfirst
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "How aRe yoU",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "THIS STRING SHOULD BE IN CAMEL CASE",
					"I got intern at geeksforgeeks":               "I got intern at geeksforgeeks",
					"Here comes the garden":                       "Here comes the garden",
					"The quick brown fox jumps over the lazy dog": "The quick brown fox jumps over the lazy dog",
				},
			},
		},
		{
			name: "hashtag",
			prepare: func(f *fields) {
				f.ChangeCaser = NewChangeCase().Hashtag
			},
			args: args{
				m: map[string]string{
					"how aRe yoU":                                 "#how #are #you",
					"THIS STRING SHOULD BE IN CAMEL CASE":         "#this #string #should #be #in #camel #case",
					"I got intern at geeksforgeeks":               "#i #got #intern #at #geeksforgeeks",
					"Here comes the garden":                       "#here #comes #the #garden",
					"The quick brown fox jumps over the lazy dog": "#the #quick #brown #fox #jumps #over #the #lazy #dog",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			for k, v := range tt.args.m {
				if got := f.ChangeCaser.Fn(k); got != v {
					t.Errorf("Case.Fn() = %v, want %v", got, v)
				}
			}
		})
	}
}
