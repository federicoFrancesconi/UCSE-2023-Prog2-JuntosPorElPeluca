const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

const urlConFiltro = "http://localhost:8080/camiones";

document.addEventListener("DOMContentLoaded", function (event) {
  if (!isUserLogged()) {
    window.location =
      document.location.origin + "/web/login/login.html?reason=login_required";
  }

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
    patente: document.getElementById("Patente").value,
    peso_maximo: parseInt(document.getElementById("PesoMaximo").value),
    fecha_creacion: "2023-10-14T12:00:00Z",
    fecha_ultima_actualizacion: "2023-10-14T12:00:00Z",
    id_creador: "",
    costo_por_kilometro: parseInt(document.getElementById("CostoPorKm").value),
  };

  console.log(JSON.stringify(data));

  makeRequest(
    `${urlConFiltro}`,
    Method.POST,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoCamion,
    errorCamion
  );
}

function exitoCamion(data) {
  window.location = window.location.origin + "/web/camiones/index.html";
}

function errorCamion(response) {
  alert(response.Error);
  console.log(response.json());
  throw new Error(response.Error);
}

function actualizarCamion(event) {
  event.preventDefault();

  const data = {
    patente: document.getElementById("Patente").value,
    peso_maximo: parseInt(document.getElementById("PesoMaximo").value),
    fecha_creacion: "2023-10-14T12:00:00Z",
    fecha_ultima_actualizacion: "2023-10-14T12:00:00Z",
    id_creador: "",
    costo_por_kilometro: parseInt(document.getElementById("CostoPorKm").value),
  };

  console.log(JSON.stringify(data));

  makeRequest(
    `${urlConFiltro}`,
    Method.PUT,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoCamion,
    errorCamion
  );

  return false;
}

function eliminarCamion(patente) {
  if (confirm("¿Estás seguro de que deseas eliminar este camión?")) {
    makeRequest(
      `${urlConFiltro}/${patente}`,
      Method.DELETE,
      null,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoCamion,
      errorCamion
    );
  } else {
    window.location = document.location.origin + "/web/camiones/index.html";
  }
}
