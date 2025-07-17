# 🛠️ Ansible Automation API

A lightweight **Go-based REST API** to **upload, list, and execute Ansible playbooks** via HTTP endpoints. Built with **Gin Web Framework** and supports safe file uploads using UUID-based naming.

---

## 🚀 Features
- List available Ansible playbooks
- Upload new playbooks securely
- Execute any uploaded playbook and return its output
- UUID prefixes prevent filename collisions

---

## 📦 Tech Stack
- Go (Golang)
- Gin Web Framework
- Ansible (installed on the system)
- UUID for unique file identification

---

## 📂 Project Structure
.
├── main.go # API Source code
├── playbooks/ # Uploaded playbooks stored here

bash
Copy
Edit

---

## 📈 API Endpoints
| Method | Endpoint                  | Description                    |
|--------|---------------------------|--------------------------------|
| GET    | `/playbooks`              | List all uploaded playbooks    |
| POST   | `/playbooks/upload`       | Upload an Ansible playbook     |
| POST   | `/playbooks/{name}/run`   | Run a specific playbook        |

Example:
```bash
curl http://localhost:8080/playbooks
curl -F "file=@your-playbook.yml" http://localhost:8080/playbooks/upload
curl -X POST http://localhost:8080/playbooks/{filename}/run
⚙️ Requirements
Go 1.18+

Ansible installed (e.g., via apt-get install ansible)

Linux/macOS recommended for Ansible execution

🏃 Getting Started
bash
Copy
Edit
git clone <repo-url>
cd ansible-automation-api
go run main.go
Access API at: http://localhost:8080

🔐 Security Recommendations
Restrict API access with authentication (JWT/API key)

Sanitize inputs to prevent command injection

Use HTTPS in production

📜 License
MIT License

yaml
Copy
Edit

---

Let me know if you want me to generate:
- **Dockerfile**
- **.env example for configurations**
- **Postman API test collection**
