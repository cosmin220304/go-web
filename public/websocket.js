// Utility class for rerendering page when server restarts
// Enabled only in development mode
if (
  window.location.hostname === "localhost" ||
  window.location.hostname === "127.0.0.1"
) {
  connectWebSocket({ reconnect: false });
}

function connectWebSocket({ reconnect }) {
  const socket = new WebSocket(`ws://${location.host}/ws`);

  socket.onopen = function (e) {
    console.log("WebSocket connection established");
    if (reconnect) {
      window.location.reload();
    }
  };

  socket.onclose = async function (event) {
    console.log(
      "WebSocket is closed. Reconnect will be attempted in 500ms.",
      event.reason
    );
    setTimeout(function () {
      connectWebSocket({ reconnect: true });
    }, 500);
  };

  socket.onerror = function (err) {
    console.error(
      "WebSocket encountered error: ",
      err.message,
      "Closing webSocket"
    );
  };
}
