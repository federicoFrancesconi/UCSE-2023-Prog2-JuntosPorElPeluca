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
  // Create an object with the data you want to send

  const now = new Date();

  const dia = String(now.getDate()).padStart(2, "0");
  const mes = String(now.getMonth() + 1).padStart(2, "0"); // Los meses comienzan desde 0
  const ano = now.getFullYear();

  const hora = String(now.getHours()).padStart(2, "0");
  const minutos = String(now.getMinutes()).padStart(2, "0");
  const segundos = String(now.getSeconds()).padStart(2, "0");

  const fechaFormateada = `${dia}/${mes}/${ano} ${hora}:${minutos}:${segundos}`;

  const data = {
    Patente: document.getElementById("Patente").value,
    PesoMaximo: parseInt(document.getElementById("PesoMaximo").value),
    FechaCreacion: "2023-10-14T12:00:00Z",
    FechaUltimaActualizacion: "2023-10-14T12:00:00Z",
    IdCreador: parseInt(document.getElementById("IdCreador").value),
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
  // Create an object with the data you want to send

  const now = new Date();

  const dia = String(now.getDate()).padStart(2, "0");
  const mes = String(now.getMonth() + 1).padStart(2, "0"); // Los meses comienzan desde 0
  const ano = now.getFullYear();

  const hora = String(now.getHours()).padStart(2, "0");
  const minutos = String(now.getMinutes()).padStart(2, "0");
  const segundos = String(now.getSeconds()).padStart(2, "0");

  const fechaFormateada = `${dia}/${mes}/${ano} ${hora}:${minutos}:${segundos}`;

  const data = {
    Patente: document.getElementById("Patente").value,
    PesoMaximo: parseInt(document.getElementById("PesoMaximo").value),
    FechaCreacion: "2023-10-14T12:00:00Z",
    FechaUltimaActualizacion: "2023-10-14T12:00:00Z",
    IdCreador: parseInt(document.getElementById("IdCreador").value),
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
    fetch(`http://localhost:8080/${patente}`, {
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
  }
}
