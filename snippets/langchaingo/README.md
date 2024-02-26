# lanchaingo

Testing out [langchaingo](https://github.com/tmc/langchaingo) paired up with [Ollama](https://github.com/ollama/ollama).

For now I froze these to go practice with the [localai](../localai/)

## Setup

```
# Download required images
docker pull ollama/ollama

# Start ollama
docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama

# Pull gemma:2b
docker exec ollama ollama run gemma:2b
```

## Proof of Concepts

1. Single Reply
2. Conversation
3. Agent Executors