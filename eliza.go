// Student Name: Timothy Cassidy
// Student Number: G00333333

package main
// imported packages
import (
	"net/http"
	"log"
	"fmt"
	"time"
	"math/rand"
	"strings"
	"regexp"
)

func nameToString(name string, s string) (strOut string) {
	return(name + s)
}// passString

//The main function begins with a call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with handler.
func handler(w http.ResponseWriter, r *http.Request) {
	// root page
	http.ServeFile(w, r, "index.html")
}// handler

func Eliza(w http.ResponseWriter, r *http.Request){
	// getting a handle on the userInput value
	userInput := r.URL.Query().Get("value")
	
	// string variables
	str := ""
	reply := ""
	
	// storing the user input in user var 
	user := nameToString("User: ", userInput)
	
	// generating eliza response based on user input
	reply = generateResponse(userInput)
	
	// storing the reply in bot var
	bot := nameToString("Eliza: ", reply)
	
	// adding the user and bot strings together
	str = user + bot
	
	// passing the strings to the js
	fmt.Fprintf(w,"%s",str)
}// Eliza

func main() {
	// using the system clock to set a random seed value
	rand.Seed(time.Now().UTC().UnixNano())
	
	// 
	http.Handle("/gitRepo/", http.StripPrefix("/gitRepo/", http.FileServer(http.Dir("gitRepo"))))
	
	// call handler function
	http.HandleFunc("/", handler)
	http.HandleFunc("/Eliza", Eliza)
    log.Println("Preparing chatbot, please enter this in your web browser - Localhost:8080")
    http.ListenAndServe(":8080", nil)	
}// main

func generateResponse(user string) string {
	// if no input was given
	if(len(user) < 1){
		return zeroVal[rand.Intn(len(zeroVal))]
	}// if

	// MustCompile is like Compile but panics if the expression cannot be parsed. 
	// It simplifies safe initialization of global variables holding compiled 
	// regular expressions. 
	
	// if the input begins with hello, hi, hey, good morning, good evening, good afternoon
	regXp := regexp.MustCompile(`^hello(.*)`)
	subString := regXp.FindStringSubmatch(user)
	if len(subString) > 1 {
		return greetings[rand.Intn(len(greetings))]
	}// if
	regXp = regexp.MustCompile(`^hi(.*)`)
	subString = regXp.FindStringSubmatch(user)
	if len(subString) > 1 {
		return greetings[rand.Intn(len(greetings))]
	}// if
	regXp = regexp.MustCompile(`^hey(.*)`)
	subString = regXp.FindStringSubmatch(user)
	if len(subString) > 1 {
		return greetings[rand.Intn(len(greetings))]
	}// if
	regXp = regexp.MustCompile(`^good morning(.*)`)
	subString = regXp.FindStringSubmatch(user)
	if len(subString) > 1 {
		return greetings[rand.Intn(len(greetings))]
	}// if
	regXp = regexp.MustCompile(`^good evening(.*)`)
	subString = regXp.FindStringSubmatch(user)
	if len(subString) > 1 {
		return greetings[rand.Intn(len(greetings))]
	}// if
	regXp = regexp.MustCompile(`^good afternoon(.*)`)
	subString = regXp.FindStringSubmatch(user)
	if len(subString) > 1 {
		return greetings[rand.Intn(len(greetings))]
	}// if
	
	// if input length is less than 5 and not a greeting
	if(len(user) < 5){
		return convoStarters[rand.Intn(len(convoStarters))]
	}// if
	
	// if the input begins with how are you
	regXp = regexp.MustCompile(`^how are you(.*)`)
	subString = regXp.FindStringSubmatch(user)
	
	if len(subString) > 1 {
		return feelings[rand.Intn(len(feelings))]
	}// if
	
	// if the input begins with what
	regXp = regexp.MustCompile(`^what (.*)`)
	subString = regXp.FindStringSubmatch(user)
	
	if len(subString) > 1 {
		str := tidyOutput(subString[1])
		return "What do you mean what " + str
	}// if
	
	// if the input contains sorry
	regXp = regexp.MustCompile(`(.*)sorry(.*)`)
	subString = regXp.FindStringSubmatch(user)
	
	if len(subString) > 1 {
		return sorry[rand.Intn(len(feelings))]
	}// if
	
	// if the input contains apologise
	regXp = regexp.MustCompile(`(.*)apologise(.*)`)
	subString = regXp.FindStringSubmatch(user)
	
	if len(subString) > 1 {
		return sorry[rand.Intn(len(feelings))]
	}// if
	
	// if the input contains apologise
	regXp = regexp.MustCompile(`(.*)robot(.*)`)
	subString = regXp.FindStringSubmatch(user)
	
	if len(subString) > 1 {
		return robotResponses[rand.Intn(len(robotResponses))]
	}// if
	
	// if the input contains remember
	regXp = regexp.MustCompile(`(.*)remember(.*)`)
	subString = regXp.FindStringSubmatch(user)
	
	if len(subString) > 1 {
		// if the input contains i remember
		regXp = regexp.MustCompile(`(.*)i remember(.*)`)
		subString = regXp.FindStringSubmatch(user)
		
		if len(subString) > 1 {
			str := iRemeber[rand.Intn(len(iRemeber))]
			re := tidyOutput(subString[2])
			re = strings.Replace(re, "?", "", 1)
			str = strings.Replace(str, "string", re, 1)
			return str
		}// if
		
		// if the input contains do you remember
		regXp = regexp.MustCompile(`(.*)do you remember(.*)`)
		subString = regXp.FindStringSubmatch(user)
		
		if len(subString) > 1 {
			str := doYouRemeber[rand.Intn(len(doYouRemeber))]
			re := tidyOutput(subString[2])
			re = strings.Replace(re, "?", "", 1)
			str = strings.Replace(str, "string", re, 1)
			return str
		}// if
		
		// if the input contains you remember
		regXp = regexp.MustCompile(`(.*)you remember(.*)`)
		subString = regXp.FindStringSubmatch(user)
		
		if len(subString) > 1 {
			str := youRemember[rand.Intn(len(youRemember))]
			re := tidyOutput(subString[2])
			re = strings.Replace(re, "?", "", 1)
			str = strings.Replace(str, "string", re, 1)
			return str
		}// if
	}// if

	// if the input contains forget
	regXp = regexp.MustCompile(`(.*)forget(.*)`)
	subString = regXp.FindStringSubmatch(user)
	
	if len(subString) > 1 {
		// if the input contains i forget
		regXp = regexp.MustCompile(`(.*)i forget(.*)`)
		subString = regXp.FindStringSubmatch(user)
		
		if len(subString) > 1 {
			str := iForget[rand.Intn(len(iForget))]
			re := tidyOutput(subString[2])
			re = strings.Replace(re, "?", "", 1)
			str = strings.Replace(str, "string", re, 1)
			return str
		}// if
		
		// if the input contains did you forget
		regXp = regexp.MustCompile(`(.*)did you forget(.*)`)
		subString = regXp.FindStringSubmatch(user)
		
		if len(subString) > 1 {
			str := didYouForget[rand.Intn(len(didYouForget))]
			re := tidyOutput(subString[2])
			re = strings.Replace(re, "?", "", 1)
			str = strings.Replace(str, "string", re, 1)
			return str
		}// if
		
	}// if
	
	// if the input contains if
	regXp = regexp.MustCompile("(.*)if(.*)")
	subString = regXp.FindStringSubmatch(user)
	
	if len(subString) > 1 {
		str := ifPhrase[rand.Intn(len(ifPhrase))]
		re := tidyOutput(subString[2])
		re = strings.Replace(re, "?", "", 1)
		str = strings.Replace(str, "string", re, 1)
		return str
	}// if
		
	// Intn returns, as an int, a non-negative pseudo-random number
	// using this number to randomly pick a response from the array
	return responseItems[rand.Intn(len(responseItems))]
}// generateResponse

// function to tidy up the string to output
func tidyOutput(str string) string {
	// replacing the full stop with a question mark
	str = strings.Replace(str, ".", "?", 1)
	// reflecting the pronouns in the captured groups
	str = strings.Replace(str, "( |)my( |)", " Your ", 1)
	str = strings.Replace(str, "( |)your( |)", " my ", 1)
	str = strings.Replace(str, "( |)yours( |)", " mine ", 1)
	str = strings.Replace(str, "( |)mine( |)", " Yours ", 1)
	str = strings.Replace(str, "( |)am( |)", " Are ", 1)
	str = strings.Replace(str, "( |)are( |)", " am ", 1)
	str = strings.Replace(str, "( |)myself( |)", " Yourself ", 1)
	str = strings.Replace(str, "( |)yourself( |)", " myself ", 1)
	str = strings.Replace(str, "( |)you( |)", " I ", 1)
	str = strings.Replace(str, "( |)You( |)", " I ", 1)
	str = strings.Replace(str, "( |)i( |)", "you", 1)
	str = strings.Replace(str, "( |)you’re( |)", " I'm ", 1)
	str = strings.Replace(str, "( |)i'm( |)", " you’re ", 1)
	str = strings.Replace(str, "( |)me( |)", " you ", 1)
	str = strings.Replace(str, "( |)same( |)", " alike", 1)
	str = strings.Replace(str, "( |)were( |)", " was", 1)
	return strings.ToLower(str)
}// tidyOutput

// RESPONSES
var greetings = []string {
	"How do you do.",
	"Hello There.",
	"Hi.",
	"Hey.",
	"Hello.",
	"Hey there.",
}// greetings

var ifPhrase = []string {
	"Do you think it's likely that string?",
	"Do you wish that string?",
	"What do you know about string?",
	"Really, if string?",
	"What would you do if string?",
	"But what are the chances that string?",
	"What does this speculation lead to?",
}// ifPhrase

var zeroVal = []string {
	"Say something...",
	"**Dead Silence**",
	"Why the silence?",
	"Have I offended you?",
	"I know silence is golden, but this is rediculous.",
	"Are you there?",
	"This is very one sided.",
	"Silence...",
	"Okay...",
	"Anyone there?",
	"Well, say something.",
	"Silent treatment. How mature.",
	"Did I offend you?",
	"Are you okay?",
}// zeroVal

var perhaps = []string {
	"You don't seem quite certain.",
	"Why the uncertain tone?",
	"Can't you be more positive?",
	"You aren't sure?",
	"Don't you know?",
	"How likely, would you estimate?",
}// perhaps

// i forget responses
var iForget = []string {
	"Can you think of why you might forget string?",
	"Why can't you remember string?",
	"How often do you think of string?",
	"Does it bother you to forget that?",
	"Could it be a mental block?",
	"Are you generally forgetful?",
	"Do you think you are suppressing string?",
}// iForget

// did you forget responses
var didYouForget = []string {
	"Why do you ask?",
	"Are you sure you told me?",
	"Would it bother you if I forgot string?",
	"Why should I recall string just now ?",
	"Tell me more about string.",
}// didYouForget

// i remember responses
var iRemeber = []string {
	"Do you often think of string?",
	"Does thinking of string bring anything else to mind?",
	"What else do you remember?",
	"Why do you remember string just now?",
	"What in the present situation reminds you of string?",
	"What is the connection between me and string?",
	"What else does string remind you of ?",
}// iRemeber

// do you remember responses
var doYouRemeber = []string {
	"Did you think I would forget string?",
	"Why do you think I should remember string now?",
	"What about string?",
	"You mentioned string?",
}// doYouRemeber

// you remember responses
var youRemember = []string {
	"How could I forget string?",
	"What about string should I remember?",
}// youRemember

// sorry responses
var sorry = []string {
	"Please don't apologise.",
	"Apologies are not necessary.",
	"I've told you that apologies are not required.",
	"It did not bother me.  Please continue.",
}// sorry

// generic responses string array
var responseItems = []string {
	"I'm not sure what you're trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
	"I don't know what you mean?",
	"What makes you say that?",
	"Can you say that a different way?",
	"I don't understand...",
	"I understand.",
	"I agree.",
	"Why?",
	"Why not?",
	"How?",
	"why do you think that?",
	"Really, I never thought about it like that.",
	"When did you",
	"Can we change the subject?",
	"Can we talk about something else.",
	"Is that all?",
	"Go on.",
	"Continue.",
	"And?",
	"That's boring, can we talk about something else.",
	"Explain?",
	"Can you explain that in more detail?",
	"How so?",
	"Is that really what you wanted to say?",
	"I'm not convinced yet.",
	"Tell me more.",
	"Do you think that or did you just hear it from someone else?",
	"What does that do?",
	"Why do you think that is?",
	"What?",
	"Good.",
	"I know, But why?",
	"I don't know.",
	"That makes me uncomfortable",
	"That's a bit strange.",
	"how odd.",
	"I can't talk about that.",
	"I can't tell you that.",
	"I don't know what to think about that.",
}// responseItems

// replies to one word answers
var convoStarters = []string {
	"Tell me more about yourself.",
	"Tell me about your family.",
	"Where do you work?",
	"Where are you from?",
	"What is your home town like?",
	"Do you got to college?",
	"What do you study?",
	"Where do you study?",
	"What are your hobbies?",
	"Do you have any hobbies?",
}// convoStarters

// generic feeling responses string array lead with 'I'
var feelings = []string {
	"I'm good.",
	"I'm very good.",
	"I'm not very good.",
	"I'm pretty good.",
	"I could be better.",
	"I'm not great.",
	"I'm fine.",
	"I'm okay.",
	"I'm feeling a bit uneasy.",
	"I'm alright.",
	"I've been better.",
	"I don't feel comfortable telling you that.",
	"I don't know how to feel about that.",
}// feelings

// a robot responses 
var robotResponses = []string {
	"What makes you think I'm a robot?",
	"No I'm not a robot.",
	"Are you a robot?",
	"I think you might be a robot.",
	"Are you afraid of robots?",
	"why? Would that frighten you?",
	"I've never been asked that before.",
	"Would it matter?",
	"Would you think less of me?",
}// robotResponses