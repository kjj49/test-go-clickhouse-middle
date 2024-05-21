-- Выборка всех уникальных eventType у которых более 1000 событий:

    SELECT eventType
    FROM events
    GROUP BY eventType
    HAVING count() > 1000;

-- Выборка событий, которые произошли в первый день каждого месяца:

    SELECT *
    FROM events
    WHERE toDate(eventTime) = toDate(toStartOfMonth(eventTime));

-- Выборка пользователей, которые совершили более 3 различных eventType:

    SELECT userID
    FROM events
    GROUP BY userID
    HAVING count(DISTINCT eventType) > 3;
