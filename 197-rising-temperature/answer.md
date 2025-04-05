# Answer
<!--
Algorithm:
- Use an alias for the table e.g. w1 and w2
- For each record w1, we find a matching row w2 where recordDate is one day before
- Then we compare their temperatures
- Only if the previous day exists, and the current temp is higher, we include the record
- This allows us to pass cases where the recorded dates may have gaps e.g. 2000-12-14, 2000-12-16; date 15 is missing hence should return nothing even if temp on 16 > 14
-->

SELECT w1.id AS Id
FROM weather w1
JOIN weather w2
  ON DATE_SUB(w1.recordDate, INTERVAL 1 DAY) = w2.recordDate
WHERE w1.temperature > w2.temperature;
