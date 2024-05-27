ARG OLLAMA_VERSION=0.1.39

FROM golang:1.21.0 as agent

RUN --mount=type=bind,source=/agent,target=/docker \
    cd /docker && \
    go build -o /ollama_agent main.go



FROM ollama/ollama:${OLLAMA_VERSION} as download_qwen_0_5b
RUN --mount=type=bind,source=/pull_model.sh,target=/pull_model.sh \
    MODEL="qwen:0.5b" ./pull_model.sh

FROM ollama/ollama:${OLLAMA_VERSION} as download_qwen_7b
RUN --mount=type=bind,source=/pull_model.sh,target=/pull_model.sh \
    MODEL="qwen:7b" ./pull_model.sh

FROM ollama/ollama:${OLLAMA_VERSION} as download_llama3_8b
RUN --mount=type=bind,source=/pull_model.sh,target=/pull_model.sh \
    MODEL="llama3:8b" ./pull_model.sh




FROM ollama/ollama:${OLLAMA_VERSION} as ollama_base
ENTRYPOINT [ "/ollama_agent" ]
EXPOSE 8000
COPY --from=agent /ollama_agent /ollama_agent
RUN mkdir -p /root/.ollama/models



FROM ollama_base as qwen_0_5b
RUN --mount=type=bind,from=download_qwen_0_5b,source=/root/.ollama/models,target=/models \
    cp -R /models/* /root/.ollama/models/

FROM ollama_base as qwen_7b
RUN --mount=type=bind,from=download_qwen_7b,source=/root/.ollama/models,target=/models \
    cp -R /models/* /root/.ollama/models/

FROM ollama_base as llama3_8b
RUN --mount=type=bind,from=download_llama3_8b,source=/root/.ollama/models,target=/models \
    cp -R /models/* /root/.ollama/models/

FROM ollama_base as multi
RUN --mount=type=bind,from=download_qwen_0_5b,source=/root/.ollama/models,target=/models \
    cp -R /models/* /root/.ollama/models/
RUN --mount=type=bind,from=download_qwen_7b,source=/root/.ollama/models,target=/models \
    cp -R /models/* /root/.ollama/models/
RUN --mount=type=bind,from=download_llama3_8b,source=/root/.ollama/models,target=/models \
    cp -R /models/* /root/.ollama/models/