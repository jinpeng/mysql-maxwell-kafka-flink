# MySQL-Maxwell-Kafka-Flink pipeline
This is an experimental pipeline that monitoring and collecting MySQL changed data via binary logs with Maxwell's Daemon, then filter and transform binlog into JSON format, then feed into Kafka, then process the data stream with Flink.

