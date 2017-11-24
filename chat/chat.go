 // Tianle Shu 
 //G00353418
 
 package chat

 import(

	 "bufio"
	 "fmt"
	// "log"
	 "math/rand"
	 "os"
	 "regexp"
	 "strings"
	 "time"
 )


 type replys struct {    //a struct called replys, contains a compiled regular expression 
	 original *regexp.Regexp
	 replacements []string      //a string array of resplacements from regexp package
 }

 func ElizaFormiles(reply string, replacements []string) replys{
	
				replys := replys{}
				original := regexp.MustCompile(reply)
				replys.original = original
				//eliza.replys = ReadReplyFormFile(replyPath)
				replys.replacements = replacements
				
				return replys
	
	
		}
	

 func ReadReplysFormFile() []replys {   //This function reads the reply.dat
  
	  entireReply := []replys{}

	  file, err := os.Open("./reply/reply.dat")
	  if err != nil {
		  panic(err)
	  }

	  defer file.Close() //the file eists

	  scanner := bufio.NewScanner(file) //read the file every line
	  
	  for scanner.Scan(){

					replyStr := scanner.Text()
					scanner.Scan()     //move to next line
					replacementsStr := scanner.Text()

					answerRow := strings.Split(replacementsStr,";")
					re := ElizaFormiles("(?!)"+replyStr, answerRow)
					entireReply = append(entireReply,re)

	  }

	  return entireReply
 }

  var reflections map[string]string    //map of strings of type string
  func Reflect(litter string) string{


	if reflections == nil { 
			
			reflections = map[string]string { 
					
					"i":      "you",
					"you":    "me",
					"me":     "you",
					"are":    "am",
					"am":     "are",
					"my":     "your",
					"i'd":    "you would",
					"i've":   "you have",
					"i'll":   "you will",
					"was":    "were",
					"you've": "I have",
					"you'll": "I will",
					"your":   "my",
					"yours":  "mine",
			
			}
					
					

		  }

		  sentences := strings.Split(litter," " )

		  for i, sentence := range sentences {

				f,s := reflections[sentence]
				if s{
						sentences[i] = f
				}
			}

				  return strings.Join(sentences," ")
		}
		  
	

  func ObtainRandonAnswer(replacements [] string) string{

			rand.Seed(time.Now().UnixNano())
			i := rand.Intn(len(replacements))
			return replacements[i]
  }   //This function make then reply sentences rand


  


 func RespondTo(text string) string {
		 responses := ReadReplysFormFile()
		 
		for _, re := range responses{

			if re.original.MatchString(text){

					output := re.original.FindStringSubmatch(text)
					
					captured := output[1]
					//boundaries := regexp.MustCompile(`[\s,.?!]+`)

					captured = Reflect(captured)
					
					replyAnswer := ObtainRandonAnswer(re.replacements)
				

					if strings.Contains(replyAnswer,"%s"){

						replyAnswer = fmt.Sprintf(replyAnswer,captured)
					}

					return replyAnswer
			}
		}

		return "I cannot understand what your mean."  
 } 




