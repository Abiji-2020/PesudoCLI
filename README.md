# PesudoCLI

Simple CLI application which uses Redis as vector store and semantic cache and helps with remembering terminal commands. Your AI-powered man pages. 

---

## Models Used
It is powered by **`Gemini AI`** for the embedding of the vector data and LLM queries. 

 - **Embedding Model:** gemini-embedding-001
 - **Chat Model:** gemini-2.0-flash 

---

# Commands

## ask 
The `ask` command is the core feature of the CLI. It answers the questions based on the vector store and the query which holds the data. And it displays the result in a specific manner. Which explains a more about the Operating System, and the command.

### Example 
```bash
pesudocli ask "Explain about podman"
```
<img width="1105" height="518" alt="image" src="https://github.com/user-attachments/assets/e37b1589-eef8-4b20-9977-c129040cc7bf" />


## init
The `init` command is used to create the vector index in the redis database. It will use the default settings such as *`HNSW`* search algorithm.

The default Index name is `pesudo_index` and the default dimension of the vector is 3076, which is one of the standard outputs from the `Gemini Embedding Model`. 

Both of these values can be updated with the flag.

### Example 

```bash
pesudocli init --index "indexname" --dim <dimenstion_value>
```

## ingest 

This command is used to embed the values from the dataset that is encoded with the binary realse. 

> [!NOTE]
> Due to the Rate Limiting we have initially have only added a 800 records in the ingest command, and they are added in the batch of 20 each. 

### Example 
```bash 
pesudocli ingest --index "indexName" --limit <limitvalue>(default=800)
```

## config 

It is a command which is used to save the configurations of the entire CLI in a safe place and to persist those values. 

It takes the arguments as :
 
- `redis-addr` : The address of the redis stack instance defaults to the local value.
- `gemini-api-key`: It is a required value and this gemini API key is used for both the embedding model and the chat model 
- `gemini-embedding-model` : If you need to update the default model to the latest or your customized one you can give it here. 
- `gemini-chat-model` : The chat LLM model which is needed for  the answering the questions.
- `index-name` : The name of the index which we will be creating in the vector Index in the Redis Instance. 

### Example 

```bash
pesudocli config --redis-addr "localhost:6379" \
--gemini-api-key "Your gemini api key" \
--gemini-embedding-model "Embedding model name" \
--gemini-chat-model "The model used for the chat" \
--index-name "Index name"
```
