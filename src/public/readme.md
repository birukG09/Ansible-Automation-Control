# Ansible Automation & Control Dashboard

A full-stack web application designed to simplify and streamline IT automation using **Ansible**. This dashboard allows users to upload, manage, and execute Ansible playbooks easily through a clean web interface — empowering faster and more reliable infrastructure management.

---

## About This Project

Automation is essential in modern IT environments to improve consistency, reduce errors, and speed up deployments. **Ansible** is a powerful open-source automation tool widely used for configuration management, application deployment, and task automation across servers.

This project bridges the gap between manual command-line Ansible usage and user-friendly automation control by providing:

- A **Golang backend** API that securely manages playbooks and executes them on demand.
- A **React frontend** that empowers users to upload Ansible playbooks, trigger runs, and instantly view execution results — no deep CLI knowledge required.

---

## Why Use Ansible?

- **Agentless:** No need to install software on target machines — it uses SSH.
- **Declarative Language:** Write human-readable YAML playbooks to define infrastructure states.
- **Extensible:** Integrates with cloud providers, networking devices, containers, and more.
- **Idempotent:** Ensures consistent results even when run multiple times.
- **Large Community & Modules:** Rich ecosystem with reusable modules for common tasks.

---

## Key Features of This Dashboard

- **Playbook Management:** Upload and store your Ansible playbooks centrally.
- **Run Automation:** Execute any uploaded playbook with a single click, automating repetitive tasks effortlessly.
- **Output Monitoring:** View detailed logs and results immediately to verify successful automation.
- **Simplified UI:** Abstracts the complexities of CLI commands behind a modern web interface.
- **Extensible Backend:** Written in Golang for fast, reliable API serving, making it easy to expand features.

---

## How It Works

1. **Upload Playbooks:** Use the frontend to upload YAML files defining your automation tasks.
2. **List & Select:** View all stored playbooks via the dashboard.
3. **Execute Playbooks:** Trigger playbook runs on the backend, which calls the native `ansible-playbook` CLI.
4. **View Results:** Execution logs and errors are returned and displayed for instant feedback.

All automation runs occur on the backend server where Ansible is installed and configured.

---

## Prerequisites

- **Go** installed (v1.18+) — for backend API  
- **Node.js & npm** — for frontend UI  
- **Ansible** installed on the backend machine — to execute playbooks via CLI  

---

## Getting Started

### Backend Setup

```bash
cd backend
go mod init ansible-backend
go get github.com/gin-gonic/gin github.com/google/uuid
go run main.go
