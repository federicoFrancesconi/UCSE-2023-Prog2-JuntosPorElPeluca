const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  if (!isUserLogged()) {
    window.location =
      document.location.origin + "/web/login/login.html?reason=login_required";
  }

  //verifico si tiene el parametro id
  const urlParams = new URLSearchParams(window.location.search);
  const idEnvio = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  if (operacion == "FINALIZAR") {
    document.getElementById("CiudadText").innerHTML = "Ciudad Final";
    document
      .getElementById("formParada")
      .addEventListener("submit", function (event) {
        finalizarViaje(idEnvio);
      });
  } else {
    document
      .getElementById("formParada")
      .addEventListener("submit", function (event) {
        agregarParada(event);
      });
  }
});

function agregarParada() {
  const urlParams = new URLSearchParams(window.location.search);
  const idEnvioParada = urlParams.get("id");

  const data = {
    id_envio: idEnvioParada,
    ciudad: document.getElementById("Ciudad").value,
    km_recorridos: parseInt(document.getElementById("KmRecorridos").value),
  };

  const urlConFiltro = `http://localhost:8080/envios/nuevaParada`;

  makeRequest(
    `${urlConFiltro}`,
    Method.POST,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoAgregarParada,
    errorAgregarParada
  );
}

function exitoAgregarParada(data) {
  window.location = document.location.origin + "/web/envios/index.html";
}

function errorAgregarParada(response) {
  alert(`Error del servidor: ${response.error}`);
  console.log(response.json());
  throw new Error(response.Error);
}

function finalizarViaje(id) {
  if (confirm("¿Estás seguro de que deseas finalizar el viaje?")) {
    //TODO: le pasamos un id como parametro a agregarParada, pero no hay funcion que reciba un id
    agregarParada();
    debugger;

    dataEnvio = {
      id: id,
      fecha_creacion: "2023-10-14T12:00:00Z",
      fecha_ultima_actualizacion: "2023-10-14T12:00:00Z",
      patente_camion: "",
      paradas: [],
      pedidos: [],
      id_creador: 0,
      estado: "Despachado",
    };

    makeRequest(
      `http://localhost:8080/envios/cambiarEstado`,
      Method.PUT,
      dataEnvio,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoAgregarParada,
      errorAgregarParada
    );
  } else {
    window.location = document.location.origin + "/web/envios/index.html";
  }
}
