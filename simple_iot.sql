CREATE TABLE "device" (
  "Id" INT(11) PRIMARY KEY NOT NULL,
  "Name" VARCHAR(50) NOT NULL,
  "ShortName" VARCHAR(3) NOT NULL,
  "Enabled" BIT(1) NOT NULL
);

CREATE TABLE "telemetry" (
  "DeviceId" INT(11) NOT NULL,
  "ReadDate" DATETIME NOT NULL,
  "Latitude" DOUBLE(10,6) DEFAULT NULL,
  "Longitude" DOUBLE(10,6) DEFAULT NULL,
  "Value" DOUBLE DEFAULT NULL,
  "Value2" DOUBLE DEFAULT NULL,
  "Value3" DOUBLE DEFAULT NULL,
  "Value4" DOUBLE DEFAULT NULL
);

ALTER TABLE "telemetry" ADD FOREIGN KEY ("DeviceId") REFERENCES "device" ("Id");

CREATE UNIQUE INDEX "UQ_Device_Name" ON "device" ("Name");

CREATE UNIQUE INDEX "UQ_Device_ShortName" ON "device" ("ShortName");

CREATE UNIQUE INDEX "UQ_ReadDate_DeviceId" ON "telemetry" ("ReadDate", "DeviceId");

CREATE INDEX "DeviceId" ON "telemetry" ("DeviceId");
