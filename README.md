# Sticky-Fingers: A Docker Clone in Go

Welcome to **Sticky-Fingers**, a lightweight Docker clone written entirely in Go! This project was created for **educational purposes**, aiming to get a better understanding of the core principles of containerization, process isolation, and how tools like Docker function under the hood.

> [!NOTE]  
> Most features listed below are not implemented yet.

---

## 🚀 Features

Sticky-Fingers implements basic containerization features, including:

- **Namespace Isolation**: Isolate process, network, and mount namespaces.
- **Control Groups (cgroups)**: Limit and prioritize resource usage (CPU, memory, etc.).
- **Image Management**: Create and manage container images.
- **Container Lifecycle Management**: Start, stop, and monitor container processes.
- **Overlay Filesystem**: Simplify image layering.

---

## 🛠️ Tech Stack

- **Programming Language**: [Go](https://golang.org/)
- **Containerization Concepts**: Namespaces, cgroups, and chroot.
- **Inspired By**: Docker's container runtime and underlying architecture.

---

## ⚠️ Disclaimer

Sticky-Fingers is built **solely for educational purposes** and is not intended for production use. It lacks the security, scalability, and reliability guarantees of tools like Docker.

---

## 🏗️ Getting Started

### Prerequisites

- **Go** (1.18 or newer)
- **Linux** (Sticky-Fingers relies on Linux-specific features like namespaces and cgroups)
- **Root Access** (Required for container isolation)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Turtel216/Sticky-Fingers.git
   cd Sticky-Fingers
   ```

2. Build the project:

   ```bash
   go build -o sticky-fingers
   ```

3. Run Sticky-Fingers:

   ```bash
   sudo ./sticky-fingers run /bin/bash
   ```

---

## 📖 How It Works

Sticky-Fingers demonstrates the core principles of containerization:

1. **Namespaces**:
   - Create isolated environments for process trees, networks, and filesystems.

2. **cgroups**:
   - Enforce resource usage limits for containers (CPU, memory, etc.).

3. **OverlayFS**:
   - Support for copy-on-write filesystems to manage image layers.

4. **Process Isolation**:
   - Run containerized processes independently of the host environment.

---

## 🛠️ Commands

| Command                     | Description                                      |
|-----------------------------|--------------------------------------------------|
| `run <command>`             | Run a command in a new container environment.    |
| `stop <container_id>`       | Stop a running container.                        |
| `ps`                        | List running containers.                         |
| `images`                    | List available container images.                 |
| `rm <container_id>`         | Remove a container.                              |

---

## 🎯 Goals of the Project

The primary goal of Sticky-Fingers is to demystify how container technologies like Docker work by providing:

- A simplified yet functional container runtime.
- Hands-on experience with Linux kernel features such as namespaces and cgroups.
- A stepping stone for developers looking to explore containerization in-depth.

---

## 🤝 Contributing

Contributions are welcome! If you find a bug or want to add a new feature, feel free to open an issue or submit a pull request.

---

## 📜 License

Sticky-Fingers is licensed under the [MIT License](LICENSE).

---

### 🌟 Acknowledgments

This project is inspired by the excellent work of Docker and other containerization tools. Special thanks to the open-source community for their invaluable resources and support.

---

### 🚧 Future Plans

- Improve image management.
- Add network isolation features.
- Create a CLI similar to Docker's for user-friendliness.

---

Enjoy exploring the magic of containerization with **Sticky-Fingers**! 🚀
