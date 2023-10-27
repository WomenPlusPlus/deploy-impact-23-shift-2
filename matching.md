# Matching algorithm idea description 

### Elements:

#### 1.Skills + Experience :
assumed predefined experience levels (beginner, intermediate, advanced), and predefined skills set
- the sum of values for every skill of job/candidate: 
  - if the skill does not match - 0
  - if the skill matches and the level is the same or the candidate level is higher than required in the job -> 1 
  - if the candidate level is lower than job required - if combination (cand: beginner job: intermediate) or (cand:intermediate job:advanced)  -> 0.7 
  - if the cand:beg job:adv -> 0.3

#### 2.Spoken Languages:
assumed predefined language levels (beginner, intermediate, advanced), and predefined language set
- the sum of values for every language of job/candidate: 
  - if the language does not match - 0
  - if the language matches and the level is the same or the candidate level is higher than required in the job -> 1 
  - if the candidate level is lower than job required - if combination (cand: beginner job:intermediate) or (cand:intermediate job:advanced)  -> 0.7
  - if the cand:beg job:adv -> 0.3

	
#### 3.Location:
(assuming that locations are predefined cities with their coordinates that enable distance count)
- measure the distance between the seek locations from the candidate and offered locations in the job and average them 

#### 4.Job location type:
(remote / office/ hybrid) 
- value depends on combinations found in the job and candidate preferences:
  - if candidate and job contain the same value -> 1
  - if existing combinations=: remote/hybrid or hybrid/office -> 0.5 
  - if no match( combination remote from one side and office from another)-> 0 

#### 5.Employment type:
(internship or long-term) 
- if match is found in employment types preferred by candidate and offered by job -> 1

#### 6.Employment rate:
(full-time or part-time) 
- if a match found in employment rate types preferred by the candidate and offered by job -> 1


#### 7.Values: 
- language model for finding value entities in both the values of the candidate and company(that posted a job) and calculating how many matched -> sum of matched

#### 8.Job description fields / Candidate about me field:  
- language model to compare similarity of candidate 'About me' field and combined job description field on the basis of recognised entities presence and synonyms match, sentiment -> similarity output value 


### Priority and weights:
1. Skills Experience -> 1
2. Spoken Languages -> 0.7
3. Location distance -> 0.6
4. Job Location type,  Employment type,  Employment rate -> 0.5
5. Values,  Job description fields / Candidate about me field  -> 0.4

###  Algorithm:
Build vector having the elements on subsequent positions for every item of the matched entity (job for candidates and candidates for jobs)
Apply weighting to positions
Calculate distances between the base vector and the matched entity vectors and choose n closest ones and return them. 
(Due to the fact that different positions of the vectors will have different value scales (some numerical and some categorical) - the matching algorithm should take into account this property)


- Companies - candidates match can be done in the same way, having all jobs posted by the company as company representation and then calculating the summed distance of vectors
