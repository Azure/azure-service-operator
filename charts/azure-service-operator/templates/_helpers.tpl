{{- define "azure-service-operator.namespace" -}}
  {{- if .Values.namespace }}
    {{- .Values.namespace }}
  {{- else }}
    {{- .Release.Namespace}}
  {{- end }}
{{- end }}
