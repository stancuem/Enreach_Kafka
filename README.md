# Kafka Data Enrichment Service
The Kafka Data Enrichment Service is a Go application designed to enhance data by retrieving messages from a Kafka topic, enriching them with additional information from a PostgreSQL database, and then writing the enriched data back to a new Kafka topic.
Key features include:
Kafka Consumer: Efficiently consumes messages from an input Kafka topic.
PostgreSQL Integration: Connects to a PostgreSQL database to fetch supplementary user information.
Data Enrichment Logic: Processes incoming data by merging it with relevant database records.
Kafka Producer: Publishes the enriched data to an output Kafka topic for further processing.
The application employs a modular architecture, making it easy to extend and maintain. It utilizes environment variables for configuration, facilitating deployment across different environments. This service demonstrates proficiency in working with Kafka, PostgreSQL, and Go, showcasing effective patterns for building data-driven applications.

