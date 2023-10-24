import { Observable, of } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { JobLocationTypeEnum, JobTypeEnum } from '@app/common/models/jobs.model';
import { Job } from '@app/jobs/common/models/job.model';

//import environment from '@envs/environment';

@Injectable({
    providedIn: 'root'
})
export class JobsService {
    constructor(private readonly httpClient: HttpClient) {
    }

    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    getJobDetails(id: number): Observable<Job> {
        // TODO: use the API.
        //return this.httpClient.get<JobModel>(`${environment.API_BASE_URL}/api/v1/jobs/${id}`);
        return of({
            title: 'Job title',
            skills: [
                'test 1',
                'test 2',
                'test 3',
                'test 4',
                'test 5',
                'test 6',
                'test 7',
                'test 8',
                'test 9',
                'test 10'
            ],
            jobType: JobTypeEnum.INTERNSHIP,
            offerSalary: 85_000,
            experienceYearFrom: 1,
            experienceYearTo: 2,
            company: {
                mission: 'To transform IT to revolutionise the enterprise.',
                name: 'ZedTech',
                values: `- Win as a team
- Innovate and execute
- Stay hungry and humble
- Deliver customer success`,
                imageUrl: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                id: 1
            },
            creator: {
                imageUrl: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                name: 'João',
                id: 2,
                email: 'joaordev@gmail.com'
            },
            creationDate: new Date(new Date().getTime() - 22 * 60 * 60 * 1000).toISOString(),
            benefits: `- Commuter benefits
- Annual learning stipends
- Work from home opportunities
- Generous family leave
- Matched donations
- Flexible PTO`,
            location: {
                city: { id: 1, name: 'Zürich' },
                type: JobLocationTypeEnum.HYBRID
            },
            candidateOverview: `- Bachelors, or Masters, degree in Design, Design Communication, Human-Computer Interaction, or related degree field
- 12 months, or less, removed from university degree program graduation date
- Experience participating in the complete product development lifecycle of web and/or software applications
- Experience in user experience design or industry experience (corporate, software, web or agency)
- An inspiring portfolio demonstrating end-to-end product design solutions, including identifying user needs, creating design solutions that address those needs, and validating solutions with end users. (Required for interview)
- A growth mindset that actively seeks learning opportunities and builds and extends core discipline skills
- Ability to build trusting relationships through clear and open communication

## Nice to have
- 1-2 years of relevant design experience (internships included)
- Experience working with design tools to create feature-level interactive designs, wireframes, and/or visual design solutions`,
            employmentLevelFrom: 80,
            employmentLevelTo: 100,
            overview:
                'ServiceNow provides cloud-based solutions that define, structure, manage, and automate services for enterprise operations, transforming old, manual ways of working into modern digital workflows. The company was founded in 2004 with a vision to build a cloud-based platform that would enable regular people to route work effectively through the enterprise.',
            rolesAndResponsibility: `- Collaborate closely with product owners and front-end developers to translate software and business requirements into actionable designs
- Support the Product team by creating UI mock-ups and prototypes to assist in requirement development
- Plan and conduct usability testing and user research sessions with internal users and clients, collecting valuable insights to inform design improvements
- This includes creating test scenarios, moderating sessions, and analyzing feedback to make data-driven design decisions
- Define and uphold the user experience standards for both new and existing product capabilities`
        });
    }
}
