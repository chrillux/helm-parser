var defaultTemplate = `apiVersion: v1
kind: Service
metadata:
  name: {{ .Chart.Name }}
  labels:
    chart: "{{ .Chart.Name }}-{{ .Chart.Version }}"
  namespace: "{{ .Release.Namespace }}"
spec:
  type: {{ .Values.service.type }}
  ports:
  - port: {{ .Values.service.externalPort }}
    targetPort: {{ .Values.service.internalPort }}
    protocol: TCP
  name: {{ .Values.service.name }}
  selector:
    app: {{ .Chart.Name }}
`;

var defaultValues = `service:
  name: nginx
  type: ClusterIP
  externalPort: 80
  internalPort: 80
`;

function clearTextAreas() {
    document.getElementById("templates").value = "";
    document.getElementById("values").value = "";
    document.getElementById("result").textContent = "";
}

function defaultTextAreas() {
    document.getElementById("templates").value = defaultTemplate;
    document.getElementById("values").value = defaultValues;
}

function renderNew(templates, values, result) {
    render(templates, values, result);
}