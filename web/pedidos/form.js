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

  const urlConFiltro = `http://localhost:8080/pedidos`;

  makeRequest(
    `${urlConFiltro}`,
    Method.POST,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoPedido,
    errorPedido
  );
}

function exitoPedido(data) {
  window.location = window.location.origin + "/web/pedidos/index.html";
}

function errorPedido(response) {
  alert("Error en la solicitud al servidor.");
  console.log(response.json());
  throw new Error("Error en la solicitud al servidor.");
}

function aceptarPedido(id) {
  if (confirm("¿Estás seguro de que deseas aceptar este pedido?")) {
    const urlConFiltro = `http://localhost:8080/${id}/aceptar`;

    makeRequest(
      `${urlConFiltro}`,
      Method.PUT,
      data,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoPedido,
      errorPedido
    );
  } else {
    window.location = document.location.origin + "/web/pedidos/index.html";
  }
}

function aceptarPedido(id) {
  if (confirm("¿Estás seguro de que deseas cancelar este pedido?")) {
    const urlConFiltro = `http://localhost:8080/${id}/cancelar`;

    makeRequest(
      `${urlConFiltro}`,
      Method.PUT,
      null,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoPedido,
      errorPedido
    );
  } else {
    window.location = document.location.origin + "/web/pedidos/index.html";
  }
}
