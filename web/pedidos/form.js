const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  //verifico si tiene el parametro id
  const urlParams = new URLSearchParams(window.location.search);
  const idPedido = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  if (idPedido != "" && idPedido != null && operacion == "ACEPTAR") {
    aceptarPedido(idPedido);
  } else if (idPedido != "" && idPedido != null && operacion == "CANCELAR") {
    cancelarPedido(idPedido);
  } else {
    document
      .getElementById("form")
      .addEventListener("submit", function (event) {
        guardarPedido(event);
      });
  }
});

function guardarPedido() {
  //armo la data a enviar
  const data = {
    id: 3,
    fecha_creacion: "2023-10-14T12:00:00Z",
    fecha_ultima_actualizacion: "2023-10-14T12:00:00Z",
    ciudad_destino: document.getElementById("CiudadDestino").value,
    productos_elegidos: [],
    id_creador: parseInt(document.getElementById("IdCreador").value),
    estado: 0,
  };

  fetch(`http://localhost:8080/pedidos`, {
    method: "POST",
    body: JSON.stringify(data),
    headers: customHeaders,
  }) // Realizar la solicitud de búsqueda (fetch) al servidor
    .then((response) => {
      if (!response.ok) {
        throw new Error("Error en la solicitud al servidor.");
      }

      window.location = "/web/pedidos/index.html";
    })
    .catch((error) => {
      console.error("Error:", error);
      alert(error);
      window.location = "/web/pedidos/form.html";
    });
}

function aceptarPedido(id) {
  if (confirm("¿Estás seguro de que deseas aceptar este pedido?")) {
    fetch(`http://localhost:8080/${id}/aceptar`, {
      method: "PUT",
      body: JSON.stringify(data),
      headers: customHeaders,
    }) // Realizar la solicitud de búsqueda (fetch) al servidor
      .then((response) => {
        if (!response.ok) {
          throw new Error("Error en la solicitud al servidor.");
        }

        window.location = "/web/pedidos/index.html";
      })
      .catch((error) => {
        console.error("Error:", error);
        alert(error);
      });
  } else {
    window.location = "/web/pedidos/index.html";
  }
}

function cancelarPedido(id) {
  if (confirm("¿Estás seguro de que deseas cancelar el pedido?")) {
    fetch(`http://localhost:8080/${id}/cancelar`, {
      method: "PUT",
      body: JSON.stringify(data),
      headers: customHeaders,
    }) // Realizar la solicitud de búsqueda (fetch) al servidor
      .then((response) => {
        if (!response.ok) {
          throw new Error("Error en la solicitud al servidor.");
        }

        window.location = "/web/pedidos/index.html";
      })
      .catch((error) => {
        console.error("Error:", error);
        alert(error);
      });
  } else {
    window.location = "/web/pedidos/index.html";
  }
}
