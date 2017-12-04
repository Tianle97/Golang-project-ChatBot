//Tianle SHU G00353418

package chat

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)


//it is a struct called reply
type reply struct {
	rex     *regexp.Regexp
	//a string array of answers as from regexp package
	answers []string
} 

//This function reads in a string array of answers and a is tring of type patter.
// It then returns an instance of reply with these loaded in.
func newReply(line string, answers []string) reply {
	reply := reply{}
	rex := regexp.MustCompile(line) //rex is the regular expression
	reply.rex = rex
	reply.answers = answers
	return reply
} //newReply



//buildReplyList reads an array of Replys from a text file.
//It takes no arguments
//buildReplyList
func buildReplyList() []reply {

	allReplys := []reply{}
	//File takes words from my reply.dat. If anything goes wrong it will exit
	file, err := os.Open("./reply/reply.dat") //open my reply.dat
	// an error
	//if err
	if err != nil {                             
		 // of err and then will escape
		panic(err)
	} 

	// The file exists!
	// this will be called by after this function.
	defer file.Close() 

	//read line by line in this file
	scanner := bufio.NewScanner(file)
	 
	//scanner
	for scanner.Scan() {

		readLine := scanner.Text()
		// move to the next line and get anwser
		scanner.Scan() 
		getAnswer := scanner.Text()
		//In reply.dat the robort replys to one input are seperated by ";"
		answerList := strings.Split(getAnswer, ";") 

		// the possible replys are "split" by the ";"
		//this regex will allow for any case (upper&lower) entered by the user
		resp := newReply("(?i)"+readLine, answerList) 

		allReplys = append(allReplys, resp)
	}

	//return the allReplys array
	return allReplys
} 


//getRandomAnswer
func getRandomAnswer(answers []string) string {
	 // seed to make it return different values.
	rand.Seed(time.Now().UnixNano())
	// Intn generates a number between 0 and num - 1
	index := rand.Intn(len(answers)) 
	return answers[index]           
}


//map of strings of type string 
var reflections map[string]string 

func Reflect(original string) string {
	//reflections from https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/
	// map hasn't been made yet
	if reflections == nil { 
		//reflections
		reflections = map[string]string{ 
			"am":     "are",
			"was":    "were",
			"i":      "you",
			"you":	  "me",
			"i'd":    "you would",
			"i've":   "you have",
			"i'll":   "you will",
			"my":     "your",
			"are":    "am",
			"myself": "yourself",
			"you've": "I have",
			"you'll": "I will",
			"your":   "my",
			"yours":  "mine",
			"me":     "you",
			"some":	  "any",
		}
	}

	words := strings.Split(original, " ")

	for index, word := range words {
		//  if it's in the map we can change the word
		val, ok := reflections[word]
		if ok { 
			words[index] = val // eg. your == mine
		}
	}

	return strings.Join(words, " ")
}

func Ask(userInput string) string {


	replys := buildReplyList()
	// look at every single reply/reply/answers
	
    //for	
	for _, resp := range replys { 
		
		//if
		if resp.rex.MatchString(userInput) {
			match := resp.rex.FindStringSubmatch(userInput)
			//match[0] is full match, match[1] is the capture group
			captured := match[1]

			captured = Reflect(captured)

			formatAnswer := getRandomAnswer(resp.answers) // get random element.

			if strings.Contains(formatAnswer, "%s") { // string needs to be formatted, %s will be sub target
				formatAnswer = fmt.Sprintf(formatAnswer, captured)
			}
			return formatAnswer

		} 

	}

	// if we're down here, it means there were no matches;
	return "Sorry I cannot understand."

}
