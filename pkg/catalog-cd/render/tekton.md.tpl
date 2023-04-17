## Workspaces

| Workspace Name    | Optional                           | Description                |
| :---------------- | :--------------------------------- | :------------------------- |
{{- range .workspaces }}
| `{{ .name }}`     | `{{ .optional | formatOptional }}` | {{ .description | chomp }} |
{{- end }}

## Params

| Param Name    | Type                       | Default                      | Description                |
| :------------ | :------------------------- | :--------------------------- | :------------------------- |
{{- range .params }}
| `{{ .name }}` | `{{ .type | formatType }}` | {{ .default | formatValue }} | {{ .description | chomp }} |
{{- end }}

## Results

| Result Name    | Description                |
| :------------- | :------------------------- |
{{- range .results }}
| `{{ .name }}`  | {{ .description | chomp }} |
{{- end }}
