import asyncio
import json
import websockets

# Set CORS headers and allowed origin (adjust origin accordingly)
allowed_origin = "*"

async def events_handler(websocket):
    # Greet the client with an initial message
    event_data = {"message": "Connection established"}
    data = json.dumps(event_data).encode()
    await websocket.send(data + b"\n\n")

    try:
        # Send events periodically
        for i in range(10):
            event_data = {"message": f"Event {i}"}
            data = json.dumps(event_data).encode()
            await websocket.send(data + b"\n\n")
            await asyncio.sleep(2)

    except websockets.ConnectionClosed:
        print("Client disconnected")

async def main():
    async with websockets.serve(events_handler, "", 8000):
        print("Server listening on port 8000")
        await asyncio.Future()

if __name__ == "__main__":
    asyncio.run(main())
