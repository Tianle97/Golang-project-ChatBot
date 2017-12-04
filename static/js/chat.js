// Creates constants for #answerList and #form id attributes from index
const answerList = $("#answerList");
const userInput = $("#userInput");


$("#sub").on("click",function(){
    res();
});

function res(){
     // form default refreshes the page and clears everything so we prevent it to make requests with JS
     event.preventDefault();
    const text = userInput.val(); // store the value taken from the form input into text
    userInput.val(""); // delete the value from the actual form input

    // if text is valid execute the following
    if (!text.trim()){
        return
    }
    
    answerList.append('<li id="right" class="list-group-item list-group-item-primary">'+ text +'<strong> :You</strong></li>');
   
    
    $.get("/chat/", {input: text})
    .done(function(w){ // "w" is the response from server
        answerList.append('<li id="left" class="list-group-item"><strong>Robort: </strong> ' + w + '</li>');
    });
}

