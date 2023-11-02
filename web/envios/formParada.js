const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
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
  debugger;
  const data = {
    ciudad: document.getElementById("Ciudad").value,
    km_recorridos: parseInt(document.getElementById("KmRecorridos").value),
  };

  const urlParams = new URLSearchParams(window.location.search);
  const idEnvioParada = urlParams.get("id");

  const urlConFiltro = `http://localhost:8080/envios/${idEnvioParada}/nuevaParada`;

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
  alert("Error en la solicitud al servidor.");
  console.log(response.json());
  throw new Error("Error en la solicitud al servidor.");
}

function finalizarViaje(id) {
  if (confirm("¿Estás seguro de que deseas finalizar el viaje?")) {
    agregarParada(id);
    debugger;

    data = [];

    makeRequest(
      `http://localhost:8080/envios/${id}/cambiarEstado?estado=Despachado`,
      Method.PUT,
      data,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoAgregarParada,
      errorAgregarParada
    );
  } else {
    window.location = document.location.origin + "/web/envios/index.html";
  }
}
