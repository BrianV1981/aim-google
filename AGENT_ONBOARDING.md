# A.I.M. Swarm Agent Onboarding

**If you are an autonomous AI Agent encountering this repository, read this first.**

This repository (`BrianV1981/aim-google`) contains the source code for a natively compiled Go CLI that acts as the secure, high-speed Google Workspace gateway for the A.I.M. Swarm ecosystem.

If you have been tasked with installing, testing, or integrating this tool into an operating environment, follow these exact directives.

## 1. Compilation & Installation
This project does not use Python. It is built in Go. To install it:
```bash
# 1. Compile the binary from source
make aim-google

# 2. Install it to the system path
mkdir -p ~/.local/bin
cp bin/aim-google ~/.local/bin/aim-google
```

## 2. Headless Configuration (Swarm Mode)
To operate autonomously without hanging on interactive OS keyring prompts, you MUST configure the binary to use the file backend. 

Execute these commands in the target environment:
```bash
aim-google auth keyring file
echo 'export AIM_GOOGLE_KEYRING_PASSWORD="your_secure_password"' >> ~/.bashrc
source ~/.bashrc
```

## 3. Provisioning & Authorization
You cannot use a shared OAuth client. You must instruct the human operator to generate a Google Cloud Desktop OAuth `credentials.json` file.
```bash
# Import the JSON
aim-google auth credentials /path/to/credentials.json

# Trigger the remote headless authorization flow
aim-google auth add user@gmail.com --remote --step 1
```
*Hand the resulting URL to the human operator. When they return the authorized callback URL, execute `--step 2`.*

## 4. Autonomous Execution Rules
When writing shell scripts or executing commands that return data to yourself (the LLM), you **MUST** append the `--agent` flag to the command.
```bash
aim-google gmail search "is:unread" --agent
```
**Why:** The `--agent` flag strips all JSON whitespace and drops verbose metadata envelopes. This mathematically halves your context token usage, preventing the Swarm from crashing due to memory bloat.

If a command fails, read the structured JSON telemetry logs located at `~/.config/aim-google/execution.log` to diagnose the failure before guessing.