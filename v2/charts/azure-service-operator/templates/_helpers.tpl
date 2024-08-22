{{/*
Create the name of the service account to use
*/}}
{{- define "azure-service-operator.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default "azureserviceoperator-default" .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
