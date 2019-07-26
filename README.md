# Go API Server

### Author: Jesse Van Volkinburg

### Setup
#### `.env` requirements
- `PORT` - Port Number for the API server to run on
- `DATABASE_URL` - PostgreSQL connection url

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
  
