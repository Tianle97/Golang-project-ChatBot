 package chat

 import(

	 "bufio"
	 "fmt"
	 "log"
	 "math/rand"
	 "os"
	 "regexp"
	 "strings"
	 "time"
 )


 type replys struct {
	 original *regexp.Regexp
	 replacements []string
 }

 func ElizaFormiles(reply string, replacements []string) replys{
	
				replys := replys{}
				original := regexp.MustCompile(reply)
				replys.original = original
				//eliza.replys = ReadReplyFormFile(replyPath)
				replys.replacements = replacements
				
				return replys
	
	
		}
	

 func ReadReplysFormFile() []replys {
  
	  entireReply := []replys{}

	  file, err := os.Open("./reply/reply.dat")
	  if err != nil {
		  log.Fatal(err)
	  }

	  defer file.Close()

	  scanner := bufio.NewScanner(file)
	  
	  for scanner.Scan(){

					replyStr := scanner.Text()
					scanner.Scan()
					replacementsStr := scanner.Text()

					answerRow := strings.Split(replyStr,";")
					resp := ElizaFormiles("(?!)"+replyStr, answerRow)
					entireReply = append(entireReply,resp)

	  }

	  return entireReply
 }

  var reflections map[string]string    //map of strings of type string
  func Reflect(litter string) string{


	if reflections == nil { 
			
			reflections = map[string]string { 
					
					"am":     "are",
					"was":    "were",
					"i":      "you",
					"i'd":    "you would",
					"i've":   "you have",
					"i'll":   "you will",
					"my":     "your",
					"are":    "am",
					"you've": "I have",
					"you'll": "I will",
					"your":   "my",
					"yours":  "mine",
					"you":    "me",
					"me":     "you",
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
  }


  



 func RespondTo(input string) string {
		 responses := ReadReplysFormFile()
		 
		for _, re := range responses{

			if matches := re.original.FileStringSubmatch(input); matches != nil{

					output := re.replacements[rand.Intn(len(re.replacements))]
					
					captured := output[1]
					//boundaries := regexp.MustCompile(`[\s,.?!]+`)

					captured = Reflect(captured)
					
					replyAnswer := ObtainRandonAnswer(re.replacements)
				

					if strings.Contains(replyAnswer,"%s"){

						replyAnswer = fmt.Split(replyAnswer,captured)
					}

					return replyAnswer
			}
		}

		return "I cannot understand what your mean."
 } 




