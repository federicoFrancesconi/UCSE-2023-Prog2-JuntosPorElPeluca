const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  //verifico si tiene el parametro id
  const urlParams = new URLSearchParams(window.location.search);
  const patente = urlParams.get("patente");
  const operacion = urlParams.get("tipo");

  if (patente != "" && patente != null && operacion == "EDITAR") {
    document
      .getElementById("form")
      .addEventListener("submit", function (event) {
        actualizarCamion(event);
      });

    document.getElementById("Patente").value = patente;
    document.getElementById("tituloFormulario").innerHTML = "Editar camion";
  } else if (patente != "" && patente != null && operacion == "ELIMINAR") {
    eliminarCamion(patente);
    document.getElementById("tituloFormulario").innerHTML = "Eliminar camion";
  } else {
    document
      .getElementById("form")
      .addEventListener("submit", function (event) {
        guardarCamion(event);
      });

    document.getElementById("tituloFormulario").innerHTML = "Crear camion";
  }
});

function guardarCamion(event) {
  event.preventDefault();

  const data = {
    Patente: document.getElementById("Patente").value,
    PesoMaximo: parseInt(document.getElementById("PesoMaximo").value),
    FechaCreacion: "2023-10-14T12:00:00Z",
    FechaUltimaActualizacion: "2023-10-14T12:00:00Z",
    IdCreador: parseInt(document.getElementById("IdCreador").value),
    CostoPorKilometro: parseInt(document.getElementById("CostoPorKm").value),
  };

  console.log(JSON.stringify(data));

  fetch(`http://localhost:8080/camiones`, {
    method: "POST",
    body: JSON.stringify(data),
    headers: customHeaders,
  }) // Realizar la solicitud de búsqueda (fetch) al servidor
    .then((response) => {
      if (!response.ok) {
        throw new Error("Error en la solicitud al servidor.");
      }

      window.location = "index.html";
    })
    .catch((error) => {
      console.error("Error:", error);
      alert(error);
    });

  return false;
}

function actualizarCamion(event) {
  event.preventDefault();

  const data = {
    patente: document.getElementById("Patente").value,
    pesoMaximo: parseInt(document.getElementById("PesoMaximo").value),
    fechaCreacion: "2023-10-14T12:00:00Z",
    fechaUltimaActualizacion: "2023-10-14T12:00:00Z",
    idCreador: parseInt(document.getElementById("IdCreador").value),
    costoPorKilometro: parseInt(document.getElementById("CostoPorKm").value),
  };

  console.log(JSON.stringify(data));

  fetch(`http://localhost:8080/camiones`, {
    method: "PUT",
    body: JSON.stringify(data),
    headers: customHeaders,
  }) // Realizar la solicitud de búsqueda (fetch) al servidor
    .then((response) => {
      if (!response.ok) {
        throw new Error("Error en la solicitud al servidor.");
      }

      window.location = "index.html";
    })
    .catch((error) => {
      console.error("Error:", error);
      alert(error);
    });

  return false;
}

function eliminarCamion(patente) {
  if (confirm("¿Estás seguro de que deseas eliminar este camión?")) {
    fetch(`http://localhost:8080/camiones/${patente}`, {
      method: "DELETE",
      headers: customHeaders,
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Error en la solicitud al servidor.");
        }
      })
      .catch((error) => {
        console.error("Error:", error);
        alert(error);
      });
  } else {
    window.location = "web/camiones/index.html";
  }
}
