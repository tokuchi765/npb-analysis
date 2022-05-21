import { rest } from '../rest';

const baseUri = '/team/stats';

const getTeamStatsByYear = async (fromYear: string, toYear: string): Promise<any> => {
  try {
    return await rest.get<any>(`${baseUri}?from_year=${fromYear}&to_year=${toYear}`);
  } catch (error: any) {
    throw new Error(error);
  }
};

export { getTeamStatsByYear };
