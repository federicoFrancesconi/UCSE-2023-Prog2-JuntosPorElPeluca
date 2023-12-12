const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  if (!isUserLogged()) {
    window.location =
      document.location.origin + "/login/login.html?reason=login_required";
  }

  //verifico si tiene el parametro id
  const urlParams = new URLSearchParams(window.location.search);
  const idEnvio = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  if (operacion == "FINALIZAR") {
    document.getElementById("CiudadText").innerHTML = "Ciudad Final";
    document
      .getElementById("buttonSave")
      .addEventListener("click", function (event) {
        finalizarViaje(idEnvio);
      });
  } else {
    document
      .getElementById("buttonSave")
      .addEventListener("click", function (event) {
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
  window.location = document.location.origin + "/envios/index.html";
}

function errorAgregarParada(status, body) {
  alert(`Error del servidor: ${body.error}`);
  console.log(body.json());
  throw new Error(status.Error);
}

function finalizarViaje(id) {
  if (confirm("¿Estás seguro de que deseas finalizar el viaje?")) {
    agregarParada();

    dataEnvio = {
      id: id,
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
    window.location = document.location.origin + "/envios/index.html";
  }
}
