# Project Configuration Files

This repository includes directories for managing YAML configuration files for projects. These files provide essential details for project setup and are intended for **development mode** and **testing**. These files should **not** be published to GitHub or any public repositories.

## Folder Structure

### 1. `project_base_info`
This folder contains YAML files that provide the basic setup information for each project. These files include general, non-sensitive information. Typical contents of the YAML files in this folder include:

- **Project Name**: The name of the project.
- **IP Address**: The base IP address for the project.
- **Port**: The port number the project runs on.

#### Example
```yaml
project_name: example_project
ip_address: 192.168.1.1
port: 8080
```

### 2. `project_personal_info`
This folder contains YAML files that store more detailed information specific to the project. These files may include sensitive information required for development and testing. Typical contents include:

- **Database Username**
- **Database Password**
- **API Keys**
- **Secrets**

#### Example
```yaml
db:
  username: admin
  password: supersecret
api:
  key: abcd1234efgh5678
secret:
  key: supersecurekey
```

## Important Notes

1. **Files Are for Development and Testing**: The configuration files in both folders are intended only for development and testing environments. Do not use these files for production purposes.

2. **Do Not Push to GitHub**: Ensure these YAML files are excluded from version control using `.gitignore` to prevent accidental publication.

3. **File Naming Convention**: Use consistent and descriptive file names for the YAML files (e.g., `project1_base.yaml`, `project1_personal.yaml`).

4. **Organized Structure**: Ensure all projects have corresponding YAML files in both folders for clarity and ease of use.

## Usage Instructions

1. Clone the repository.
2. Navigate to the respective folder based on the type of information you need.
3. Use the contents of the YAML files to configure your project environment for development or testing.

By maintaining a clear distinction between base information and personal information, and adhering to these guidelines, you can streamline your project's configuration process effectively.

