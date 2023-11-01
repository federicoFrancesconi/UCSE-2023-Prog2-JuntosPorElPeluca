document.addEventListener("DOMContentLoaded", function (event) {
  //verifico si tiene el parametro id
  const urlParams = new URLSearchParams(window.location.search);
  const idEnvio = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  if (idEnvio != "" && idEnvio != null && operacion == "INICIAR") {
    iniciarViaje(idEnvio);
  } else if (idEnvio != "" && idEnvio != null && operacion == "FINALIZAR") {
    finalizarViaje(idEnvio);
  } else {
    document
      .getElementById("form")
      .addEventListener("submit", function (event) {
        guardarEnvio(event);
      });
  }
});

const urlConFiltro = `http://localhost:8080/envios`;

function guardarEnvio() {
  //obtengo los datos de los pedidos
  const valorPedidos = document.getElementById("Pedidos").value;
  const valoresSeparados = valorPedidos.split(",");
  const pedidosArray = valoresSeparados.map(function (valor) {
    return parseInt(valor);
  });

  //armo la data a enviar
  const data = {
    id: 3,
    fecha_creacion: "2023-10-14T12:00:00Z",
    fecha_ultima_actualizacion: "2023-10-14T12:00:00Z",
    patente_camion: document.getElementById("PatenteCamion").value,
    paradas: [],
    pedidos: pedidosArray,
    id_creador: parseInt(document.getElementById("IdCreador").value),
    estado: "ADespachar",
  };

  makeRequest(
    `${urlConFiltro}`,
    Method.POST,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoEnvio,
    errorEnvio
  );
}

function exitoEnvio(data) {
  window.location = window.location.origin + "/web/envios/index.html";
}

function errorEnvio(response) {
  alert("Error en la solicitud al servidor.");
  console.log(response.json());
  throw new Error("Error en la solicitud al servidor.");
}

function iniciarViaje(id) {
  if (confirm("¿Estás seguro de que deseas iniciar el viaje?")) {
    makeRequest(
      `${urlConFiltro}/${id}/cambiarEstado?estado=EnRuta`,
      Method.PUT,
      data,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoEnvio,
      errorEnvio
    );
  } else {
    window.location = document.location.origin + "/web/envios/index.html";
  }
}

function finalizarViaje(id) {
  if (confirm("¿Estás seguro de que deseas finalizar el viaje?")) {
    makeRequest(
      `${urlConFiltro}/${id}/cambiarEstado?estado=Despachado`,
      Method.PUT,
      data,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoEnvio,
      errorEnvio
    );
  } else {
    window.location = document.location.origin + "/web/envios/index.html";
  }
}
