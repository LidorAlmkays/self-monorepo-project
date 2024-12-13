# Project Environment Configuration

This repository includes two directories for managing environment configuration files for projects. These files provide essential details for project setup but are excluded from version control for security and privacy reasons.

## Folder Structure

### 1. `project_base_info`
This folder contains environment files that provide the basic setup information for each project. These files are shared across the team but exclude any sensitive details. Typical contents of the environment files in this folder include:

- **Project Name**: The name of the project.
- **IP Address**: The base IP address for the project.
- **Port**: The port number the project runs on.

#### Example
```env
PROJECT_NAME=example_project
IP_ADDRESS=192.168.1.1
PORT=8080
```

### 2. `project_personal_info`
This folder contains environment files that store private and sensitive information specific to the project. These files are not shared publicly and must be managed securely. Typical contents include:

- **Database Username**
- **Database Password**
- **API Keys**
- **Secrets**

#### Example
```env
DB_USERNAME=admin
DB_PASSWORD=supersecret
API_KEY=abcd1234efgh5678
SECRET_KEY=supersecurekey
```

## Important Notes

1. **Do Not Push to GitHub**: Both `project_base_info` and `project_personal_info` folders are excluded from version control using `.gitignore`. Ensure sensitive information is never pushed to a public or shared repository.

2. **File Naming Convention**: Use consistent and descriptive file names for the environment files (e.g., `project1.env`, `project2.env`).

3. **Environment Variables**: Always load these environment files securely in your project (e.g., using libraries like `dotenv` in Node.js or `os` in Python).

4. **Access Management**: Limit access to the `project_personal_info` folder to authorized personnel only.

## Usage Instructions

1. Clone the repository.
2. Navigate to the respective folder based on the type of information you need.
3. Use the contents of the `.env` files to configure your project environment.
4. Never share or expose the contents of the `project_personal_info` folder.

## Security Best Practices

- **Encrypt Sensitive Files**: Use encryption tools to secure sensitive environment files.
- **Rotate Keys and Passwords**: Regularly update sensitive keys and passwords.
- **Use Secure Channels**: Share private environment files only through secure channels.

By following the above practices, you can ensure that your project's environment configuration remains secure and manageable.

