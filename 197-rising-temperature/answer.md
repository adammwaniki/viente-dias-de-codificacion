# Answer
<!--
Algorithm:
- Use a WITH and a common table expression
    - A common table expression (CTE) is a named temporary result set that exists within the scope of a single statement and that can be referred to later within that statement, possibly multiple times
- Inside this CTE we can obtain the temperature and use LAG or LEAD to compare the current temperature value to either the previous temperature value or the next temperature value respectively
- recordDate is a unique and ordered value ,therefore, we can ORDER BY it
- else we could just order by id if autoincrementing
-->

WITH CTE AS (
  SELECT
    id,
    temperature,
    LAG(temperature) OVER (ORDER BY recordDate) AS prev_temperature
  FROM weather
)
SELECT
  id
FROM CTE
WHERE temperature > prev_temperature;
