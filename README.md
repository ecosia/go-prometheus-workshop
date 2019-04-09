# Workshop: Monitoring Go Applications with Prometheus

## Objective

In the directory `app/`, we have a simple Go application. We want to start observing the behaviour of this application at runtime, by tracking and exporting metric data.

We will do this using the time-series database system [Prometheus](https://prometheus.io), which uses a "pull" method to extract data from running applications. This means that the applications need to "export" their data, so that Prometheus is able to "scrape" the metric data from them. This is typically done via an HTTP endpoint (`/metrics`, by convention).

We will use the [Prometheus Go client library](https://godoc.org/github.com/prometheus/client_golang/prometheus) to track metrics in our code, and to export the metrics on `/metrics`.

## Content

### Section 1: Exporting metrics

Exporting basics:

Read the documentation or examples about the Prometheus Go client. In particular, you can check the [simple example]() which demonstrates usage of the `promhttp` - this includes a `.Handler()` function which returns an `http.Handler`. The Prometheus Go client exports many metrics by default (about the Go runtime, eg. garbage collection), so you can export just these default metrics by simply attaching the `promhttp` handler to an `http.Server`. For example, if you have a muxer in a variable called `mux`, you can call `mux.Handle("/metrics", promhttp.Handler())`. You should then be able to start the server, and see some default metrics being exported on `/metrics`.


Custom metrics:

Then, you'll want to track and export custom metrics. This is a three-step process: creating a metric (of a given data type); registering the metric; tracking the metric.

Prometheus has a few different data types, but the simplest is a `Counter` - this is a counter which always goes up, and can be used to track, for example, the number of requests received (you can then divide this unit over time to calculate requests per second). To create a `Counter`, you can use the `prometheus` Go client, with `prometheus.NewCounter(opts)`, where `opts` is a `prometheus.CounterOpts` (a struct containing metadata - at minimum, a `Name`). You can store this in a variable, like:

    requestCounter := prometheus.NewCounter(prometheus.CounterOpts{Name: "requests_total"})

After creating a metric, you still won't see it appear in `/metrics` until it's been "registered". You can do this with `prometheus.MustRegister(metric)`, which will attempt to register the metric and panic if it fails (the non-panicing version also exists, as `prometheus.Register()`, but for this workshop, we recommend using `MustRegister()`). Then, you should be able to see your metric exposed on `/metrics` - success! (Except, it will still always report 0 - not quite useful, yet)

To use our metric in practice, we want to increment the counter when tracking events in our code. To increment the `Counter` type by one, we can simply call `.Inc()` - for example, using the request counter we created above, we could call:

    requestCounter.Inc()

You should add these `.Inc()` calls in the place in your code where the event you want to track is occuring.


Labels:

Labels are a way of adding contextual information to your metrics (increasing their "cardinality"). For example, when tracking the count of requests received, it might be useful to also track the status code of the request. To do this, you can use `prometheus.NewCounterVec()` instead of `prometheus.NewCounter()`, and provide a list of label keys as the second argument - for example:

    requestCounter := prometheus.NewCounter(prometheus.CounterOpts{Name: "requests_total"}, []string{"status_code"})

Then, when tracking it, you'll need to provide the label values, which can be done like so:

    var status int
    if something.wasSuccessful:
        status = http.StatusOK
    else:
        status = http.StatusInternalServerError

	requestCounter.WithLabelValues(status).Inc()

