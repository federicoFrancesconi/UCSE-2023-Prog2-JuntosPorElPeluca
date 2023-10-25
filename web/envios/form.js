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
    estado: 0,
  };

  fetch(`http://localhost:8080/envios`, {
    method: "POST",
    body: JSON.stringify(data),
    headers: customHeaders,
  }) // Realizar la solicitud de búsqueda (fetch) al servidor
    .then((response) => {
      if (!response.ok) {
        throw new Error("Error en la solicitud al servidor.");
      }

      window.location = "/web/envios/index.html";
    })
    .catch((error) => {
      console.error("Error:", error);
      alert(error);
      window.location = "/web/envios/form.html";
    });
}

function iniciarViaje(id) {
  if (confirm("¿Estás seguro de que deseas iniciar el viaje?")) {
    fetch(`http://localhost:8080/${id}/iniciar`, {
      method: "PUT",
      body: JSON.stringify(data),
      headers: customHeaders,
    }) // Realizar la solicitud de búsqueda (fetch) al servidor
      .then((response) => {
        if (!response.ok) {
          throw new Error("Error en la solicitud al servidor.");
        }

        window.location = "/web/envios/index.html";
      })
      .catch((error) => {
        console.error("Error:", error);
        alert(error);
      });
  } else {
    window.location = "/web/envios/index.html";
  }
}

function finalizarViaje(id) {
  if (confirm("¿Estás seguro de que deseas finalizar el viaje?")) {
    fetch(`http://localhost:8080/${id}/finalizar`, {
      method: "PUT",
      body: JSON.stringify(data),
      headers: customHeaders,
    }) // Realizar la solicitud de búsqueda (fetch) al servidor
      .then((response) => {
        if (!response.ok) {
          throw new Error("Error en la solicitud al servidor.");
        }

        window.location = "/web/envios/index.html";
      })
      .catch((error) => {
        console.error("Error:", error);
        alert(error);
      });
  } else {
    window.location = "/web/envios/index.html";
  }
}
