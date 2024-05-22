## Notes

when using the hash table data structure, you may likely experience a situation where the hash function generates he same hash code for  two different keys. 

To handle this, you use the collision handling techniques which includes 
- open addressing 
- separate chaining. 

Open addressing searches for the next available slot within  the table when the two keys are having one spot. Separate chaining handles collisions by storing multiple keys in a single slot using a linked list