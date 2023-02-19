# ChatGenie

ChatGenie is a command-line interface (CLI) based chatbot that utilizes the GPT-3 language model from OpenAI to generate natural and human-like responses to user input. It allows users to interact with the chatbot through the command-line interface, making it easy to use and accessible. ChatGenie uses the power of GPT-3 to understand the user's input and generate relevant and intelligent responses. The chatbot can be used for various purposes like customer service, knowledge base, and even for entertainment purposes. With ChatGenie, you can have a conversation with the chatbot just like you would with a human, making it a fun and interactive experience.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You will need the following to run the project:

-   Go (1.15 or higher)
-   A OpenAI API key
-   Add the OpenAI api_key in .env file

### Setting Up the CLI

1.  Clone the repository to your local machine:

```
git clone https://github.com/yrs147/go-gpt-cli.git
```

2.  Navigate to the project directory:

```
cd ChatGenie
```
3. Install Dependencies

```
go mod tidy
```

4.  Generate an  API key from the OPEN AI [website](https://beta.openai.com/account/api-keys)

![image](https://user-images.githubusercontent.com/98258627/213876014-b3f15700-e279-4515-bb25-b706d63e75b6.png)

5. Add it in the .env file 

![image](https://user-images.githubusercontent.com/98258627/213876103-4b5ded86-98ee-461e-bf14-5dbd09bdc7fc.png)

6. Run the project 

```
go run main.go

```
7.  Start chatting with Chatbot by typing in a question or statement and pressing enter.

### Using the chatbot


![image](https://user-images.githubusercontent.com/98258627/213878247-fe13a9b2-4ae0-4ab0-b906-4470495c0dab.png)

The chatbot will continue running until you type in **quit** and press enter.

