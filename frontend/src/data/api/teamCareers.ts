import { rest } from '../rest';

const baseUri = '/team/careers';

const getCareers = async (teamID: string | undefined, year: string): Promise<any> => {
  try {
    return await rest.get<any>(`${baseUri}/${teamID}/${year}`);
  } catch (error: any) {
    throw new Error(error);
  }
};

export { getCareers };
