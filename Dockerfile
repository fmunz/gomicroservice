# initial image, nada
FROM scratch 

COPY fmhttp /

# port used by fmhttp
EXPOSE 5555
# default command to execute (only default!)
CMD ["/fmhttp"]
