# Go API Server

### Author: Jesse Van Volkinburg

### Setup
#### `.env` requirements
- `PORT` - Port Number for the API server to run on
- `DATABASE_URL` - PostgreSQL connection url to your database

#### Initializing the database
- Create a new database in your PostgreSQL server
- Set the `DATABASE_URL` environment variable or add it to the `.env` file in the same directory as the server binary
- `./go-apiserver` will start the API server
- Perform a `GET` request at the route endpoint `/database/initialize`
  This can be performed from any standard HTTP client, including a browser.
- Server will respond with a HTTP Status 200 and a message `Database Initialized`

#### Running the app
- `./go-apiserver` will start the server
- Endpoint: `/categories`
  - `GET`: Returns a JSON object with a list of categories
  - `POST`: Add a category to the database. Send a request body, encoded as `application/json`, of the form:
    ```
    {
      name: "name",
      display_name: "display name",
      description: "description"
    }
    ```
    Return value will be the added category, with id, as a JSON object.
- Endpoint: `/categories/:id`
  - `GET`: Returns a JSON object with the category requested by the given id
  - `DELETE`: Deletes a category specified by the given id from the database, if it exists. Returns a JSON object with the deleted category in it.
  
