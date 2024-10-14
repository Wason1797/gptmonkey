# GPTmonkey

This is a simple command that talks with codellama to ask for commands.

## Usage

Add your configuration in `.gptmonkey` like

```text
CODELLAMA_URL=http://localhost:11434/api/generate
```
Call it with

```bash
gptmonkey ssh into a server with ip 10.10.10.1
```
<details>
  <summary>Output:</summary>

To SSH into a server with the IP address 10.10.10.1, you can use the following command:
```
ssh username@10.10.10.1
```
Replace `username` with your actual username for the server. If you are using a password to authenticate, you will be prompted to enter it when you run this command. If you have set up key-based authentication, you may need to specify the path to your private key file using the `-i` option, like this:
```
ssh -i /path/to/private/key username@10.10.10.1
```
Note that you will need to replace `/path/to/private/key` with the actual path to your private key file on your local system

</details>

## Requirements

You need to install ollama and codellama. This model runs locally, or you can add a remote ollama url.

## TODOS:

- [ ] Add a configuration to only output code
- [ ] Add a loading screen while the model is thinking
- [ ] Add propper instalation instructions
- [ ] Add colors
- [ ] Install with brew
