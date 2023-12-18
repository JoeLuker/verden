{{/*
Expand the name of the chart.
*/}}
{{- define "verden-chart.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default label.
*/}}
{{- define "verden-chart.labels" -}}
helm.sh/chart: {{ include "verden-chart.chart" . }}
{{ include "verden-chart.selectorLabels" . }}
{{- end -}}

{{/*
Selector labels.
*/}}
{{- define "verden-chart.selectorLabels" -}}
app.kubernetes.io/name: {{ include "verden-chart.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "verden-chart.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}