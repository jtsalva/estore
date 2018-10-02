CREATE TABLE `Tags` (
    `TagId` INT NOT NULL AUTO_INCREMENT,
    `Name` VARBINARY(40) NOT NULL UNIQUE,
    PRIMARY KEY (`TagId`)
);

CREATE TABLE `Categories` (
    `CategoryId` INT NOT NULL AUTO_INCREMENT,
    `Name` VARBINARY(40) NOT NULL UNIQUE,
    PRIMARY KEY (`CategoryId`)
);

CREATE TABLE `Roles` (
    `RoleId` INT NOT NULL AUTO_INCREMENT,
    `Name` VARBINARY(60) NOT NULL UNIQUE,
    PRIMARY KEY (`RoleId`)
);

CREATE TABLE `Users` (
    `UserId` INT NOT NULL AUTO_INCREMENT,
    `Name` VARBINARY(60) NOT NULL,
    `Email` VARBINARY(320) NOT NULL UNIQUE,
    `Password` BINARY(89) NOT NULL,
    `DateJoined` TIMESTAMP NOT NULL,
    `Role` INT NOT NULL,
    PRIMARY KEY (`UserId`),
    FOREIGN KEY (`Role`) REFERENCES Roles(`RoleId`)
);

CREATE TABLE `Items` (
    `ItemId` INT NOT NULL AUTO_INCREMENT,
    `Name` VARBINARY(75) NOT NULL,
    `Description` VARBINARY(8000) DEFAULT "",
    `Price` DECIMAL(13, 4) NOT NULL,
    `Category` INT DEFAULT 1,
    PRIMARY KEY (`ItemId`),
    FOREIGN KEY (`Category`) REFERENCES Categories(`CategoryId`)
);

CREATE TABLE `Orders` (
    `OrderId` INT NOT NULL AUTO_INCREMENT,
    `Customer` INT NOT NULL,
    PRIMARY KEY (`OrderId`),
    FOREIGN KEY (`Customer`) REFERENCES Users(`UserId`)
);

CREATE TABLE `SoldItems` (
    `ItemId` INT NOT NULL,
    `OrderId` INT NOT NULL,
    `Quantity` SMALLINT NOT NULL,
    PRIMARY KEY (`ItemId`, `OrderId`),
    FOREIGN KEY (`ItemId`) REFERENCES Items(`ItemId`),
    FOREIGN KEY (`OrderId`) REFERENCES Orders(`OrderId`)
);

CREATE TABLE `ItemTags` (
    `ItemId` INT NOT NULL,
    `TagId` INT NOT NULL,
    PRIMARY KEY (`ItemId`, `TagId`),
    FOREIGN KEY (`ItemId`) REFERENCES Items(`ItemId`),
    FOREIGN KEY (`TagId`) REFERENCES Tags(`TagId`)
);