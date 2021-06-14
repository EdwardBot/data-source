FROM centurylink/ca-certs
ADD main /
ENV PORT=3000
EXPOSE ${PORT}
CMD ["/main"]
