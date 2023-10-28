const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  //verifico si tiene el parametro id
  const urlParams = new URLSearchParams(window.location.search);
  const codProducto = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  if (codProducto != "" && codProducto != null && operacion == "ELIMINAR") {
    eliminarProducto(codProducto);
  } else {
    document
      .getElementById("form")
      .addEventListener("submit", function (event) {
        guardarProducto(event);
      });
  }
});

function guardarProducto() {
  //armo la data a enviar
  const data = {
    codigo_producto: document.getElementById("CodigoProducto").value,
    fecha_creacion: "2023-10-14T12:00:00Z",
    fecha_ultima_actualizacion: "2023-10-14T12:00:00Z",
    tipo_producto: document.getElementById("TipoProducto").value,
    nombre: document.getElementById("Nombre").value,
    peso_unitario: parseFloat(document.getElementById("PesoUnitario").value),
    precio_unitario: parseFloat(
      document.getElementById("PrecioUnitario").value
    ),
    stock_minimo: parseInt(document.getElementById("StockMinimo").value),
    stock_actual: parseInt(document.getElementById("StockActual").value),
    id_creador: parseInt(document.getElementById("IdCreador").value),
  };

  fetch(`http://localhost:8080/productos`, {
    method: "POST",
    body: JSON.stringify(data),
    headers: customHeaders,
  }) // Realizar la solicitud de búsqueda (fetch) al servidor
    .then((response) => {
      if (!response.ok) {
        throw new Error("Error en la solicitud al servidor.");
      }

      window.location = "/web/productos/index.html";
    })
    .catch((error) => {
      console.error("Error:", error);
      alert(error);
      window.location = "/web/productos/form.html";
    });
}

function eliminarProducto(codProducto) {
  if (confirm("¿Estás seguro de que deseas eliminar este producto?")) {
    fetch(`http://localhost:8080/productos/${codProducto}`, {
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
    window.location = "web/productos/index.html";
  }
}
