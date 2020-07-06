-- Find all subsector capitals
select id, name, subsector_index from world where sector_id = 216 and remarks like '%Cp%' 

-- More detailed find all subsector capitals
select subsector.id, subsector.name, subsector.sector_index, subsector.capital_id, world.id, world.name from subsector,world where 
world.sector_id=216 and subsector.sector_id=216 and world.subsector_index = subsector.sector_index and world.remarks like '%Cp%'

-- Updates Subsector info (subsector id = 216) with subsector capital information 
UPDATE subsector
  SET capital_id = (
	SELECT id FROM world WHERE world.sector_id=216 AND world.remarks like '%Cp%' AND world.subsector_index=subsector.sector_index)
	

-- Search for duplicates in the sectors table
SELECT x_loc, y_loc, count(*)
FROM sector
GROUP BY x_loc, y_loc
HAVING COUNT(*) > 1;


-- Name those duplicates
SELECT a.*
FROM sector a
JOIN (SELECT x_loc, y_loc, count(*)
FROM sector
GROUP BY x_loc, y_loc
HAVING COUNT(*) > 1 ) b
ON a.x_loc = b.x_loc
AND a.y_loc = b.y_loc
ORDER by a.name; 