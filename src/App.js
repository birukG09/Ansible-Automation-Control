import React, { useEffect, useState } from "react";
import axios from "axios";

function App() {
  const [playbooks, setPlaybooks] = useState([]);
  const [selectedFile, setSelectedFile] = useState(null);
  const [output, setOutput] = useState("");
  const backendUrl = "http://localhost:8080";

  const fetchPlaybooks = async () => {
    try {
      const res = await axios.get(`${backendUrl}/playbooks`);
      setPlaybooks(res.data.playbooks);
    } catch (err) {
      alert("Failed to load playbooks");
    }
  };

  useEffect(() => {
    fetchPlaybooks();
  }, []);

  const onFileChange = (e) => {
    setSelectedFile(e.target.files[0]);
  };

  const uploadPlaybook = async () => {
    if (!selectedFile) return alert("Select a file first");
    const formData = new FormData();
    formData.append("file", selectedFile);
    try {
      await axios.post(`${backendUrl}/playbooks/upload`, formData);
      setSelectedFile(null);
      fetchPlaybooks();
      alert("Uploaded successfully");
    } catch (err) {
      alert("Upload failed");
    }
  };

  const runPlaybook = async (name) => {
    setOutput("Running...");
    try {
      const res = await axios.post(`${backendUrl}/playbooks/${name}/run`);
      setOutput(res.data.output);
    } catch (err) {
      setOutput(err.response?.data?.output || "Run failed");
    }
  };

  return (
    <div style={{ maxWidth: 700, margin: "auto", padding: 20 }}>
      <h1>Ansible Automation Dashboard</h1>

      <div>
        <input type="file" accept=".yml,.yaml" onChange={onFileChange} />
        <button onClick={uploadPlaybook}>Upload Playbook</button>
      </div>

      <h2>Playbooks</h2>
      <ul>
        {playbooks.map((pb) => (
          <li key={pb}>
            {pb}{" "}
            <button onClick={() => runPlaybook(pb)}>Run</button>
          </li>
        ))}
      </ul>

      <h2>Output</h2>
      <pre style={{ whiteSpace: "pre-wrap", background: "#f0f0f0", padding: 10 }}>
        {output}
      </pre>
    </div>
  );
}

export default App;
