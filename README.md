<h1>Go Auth with JWT</h1>
<p>A fast, lightweight, and secure JWT authentication service built using <strong>Go</strong>, <strong>MySQL</strong>, and <strong>Docker</strong>
</p>
<p>Ideal for developers who want a quick and secure way to implement user authentication in their apps. With Docker containerization, this service is also easily deployable and scalable for modern applications.</p>
<h2>⚙️ Tech Stack</h2>
<ul>
  <li>
    <strong>GO:1.23.4</strong> (Gofiber v2, GORM)
  </li>
  <li>
    <strong>MySQL:8.0.30</strong> (Database)
  </li>
  <li>
    <strong>Docker</strong> (Containerization)
  </li>
</ul>
<h2>📦 Getting Started</h2>
<h3>Prerequisites</h3>
<ul>
  <li>
    <strong>Docker</strong>: Make sure Docker is installed and running.
  </li>
  <li>
    <strong>MySQL</strong>: MySQL should be configured (will run in a Docker container).
  </li>
</ul>
<h3>Clone the repository</h3>
<pre>
  git clone https://github.com/your-username/go-auth-jwt.git
  cd go-auth-jwt
</pre>
<h3>Docker Setup</h3>
<ol>
  <li>
    <strong>Build the Docker container</strong>:
    <pre>
  docker build --tag go_auth_jwt:1.0 .
</pre>
  </li>
  <li>
    <strong>Start the application</strong>:
    <pre class="!overflow-visible">
  docker compose up -d
</pre>
  </li>
  <li>The app will now be accessible on <code>http://localhost:9000</code>. </li>
</ol>
<h3>Database Configuration</h3>
<ul>
  <li>The MySQL database will run in a Docker container and will be automatically configured on startup.</li>
  <li>You might migrate your DB by accessing the MySQL database on <code>localhost:3306</code> with the credentials set in the <code>docker-compose.yml</code> file. </li>
</ul>
<h2>🌐 API Endpoints</h2>


```
###
POST http://localhost:9000/signup HTTP/1.1
Content-Type: application/json
{
    "name": "fulan",
    "email": "ful@n.com",
    "password": "123456"
}

###
POST http://localhost:9000/login HTTP/1.1
Content-Type: application/json

{
    "email": "ful@n.com",
    "password": "123456"
}

###
GET http://localhost:9000/private HTTP/1.1
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY5MDUxMTgsInVzZXJfaWQiOjE0fQ.CSH6gkLwnDVVa_9bYk43bOTFnTxPSm6c6OeKzR8sg_s
```
