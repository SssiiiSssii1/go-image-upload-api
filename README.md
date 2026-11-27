# go-image-upload-api

A small image upload and serving API written in Go with [Gin](https://github.com/gin-gonic/gin).

## Features

- Upload an image with a multipart form request
- Accepts JPEG, PNG and WebP only
- Rejects files larger than 5 MB
- Stores each file under a random UUID filename, keeping the original extension
- Serves stored images by filename

## Requirements

- Go 1.25 or newer

## Run

    go run .

The server listens on port 3000. Uploaded files are written to the `uploads/` directory.

## API

### Upload an image

    POST /images

Send a multipart form with the file in the `img` field:

    curl -F "img=@photo.png" http://localhost:3000/images

Response `201`:

    {"filename": "<uuid>.png", "size": 12345}

Errors return `400` with an `error` message (missing image, unsupported type, file too large) or `500` if saving fails.

### Get an image

    GET /images/:filename

    curl http://localhost:3000/images/<uuid>.png
