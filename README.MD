# Calculator API

This is a basic RESTful API built with Go that provides endpoints for simple arithmetic calculations: addition, subtraction, multiplication, and division. 
It also includes a welcome message on the root endpoint.

## Endpoints

* **`/` (GET):** Returns a welcome message.
* **`/add` (POST):** Adds two numbers provided in the request body.
* **`/subtract` (POST):** Subtracts two numbers provided in the request body.
* **`/multiply` (POST):** Multiplies two numbers provided in the request body.
* **`/divide` (POST):** Divides two numbers provided in the request body.

## Request and Response Format

The API expects input numbers for the calculation endpoints to be provided in the JSON request body with the keys `"a"` and `"b"`. The values should be integers.

The API responds to successful calculation requests with a JSON body containing the result under the key `"result"`.

## Example Usage: `/add` Endpoint

This section demonstrates how to use the `/add` endpoint to add two numbers.

**Request:**

* **Method:** `POST`
* **Endpoint:** `/add`
* **Headers:** `Content-Type: application/json`
* **Body (JSON):**

    ```json
    {
        "a": 10,
        "b": 5
    }
    ```

**Example using `curl` (command line):**

```bash
curl -X POST -H "Content-Type: application/json" -d '{"a": 10, "b": 5}' http://your-api-domain:your-api-port/add