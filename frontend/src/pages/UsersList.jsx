import React, {useEffect, useState} from 'react';
import {useQuery} from '@tanstack/react-query';
import {fetchUsers} from '../services/statsService.js'; // replace with your actual service

const UsersList = () => {
    const [currentPage, setCurrentPage] = useState(1);
    const [itemsPerPage] = useState(10);
    const [totalItem, setTotalItem] = useState(0);

    const {data, isLoading, error} = useQuery({
        queryKey: ['users=', currentPage, itemsPerPage],
        queryFn: () => fetchUsers(currentPage, itemsPerPage)
    });

    useEffect(() => {
        if (currentPage === 1) {
            setTotalItem(data?.totalItems ?? 0)
        }
    }, [data?.totalItems, currentPage, itemsPerPage]);

    const users = data?.data ?? []

    if (error) return <p>Error loading user data</p>;

    const handlePageChange = (pageNumber) => {
        if (pageNumber > Math.ceil(totalItem / itemsPerPage) || pageNumber < 1) return;
        setCurrentPage(pageNumber);
    };

    return (
        <div className="mt-10 vh max-w-screen-2xl mx-auto px-4">
            <h1 className="text-gray-100 text-3xl font-bold text-center">Users</h1>

            <div className="overflow-x-auto">

                <div className="flex justify-between">
                    <h4>Total users: ({totalItem})</h4>
                    <div className="pagination flex justify-center py-4 gap-x-2">
                        <button
                            className="btn btn-primary btn-sm"
                            onClick={() => handlePageChange(currentPage - 1)}
                            disabled={currentPage === 1}
                        >
                            Prev
                        </button>
                        <span className="btn btn-sm">
          Page {currentPage}
        </span>
                        <button
                            className="btn btn-primary btn-sm"
                            onClick={() => handlePageChange(currentPage + 1)}
                            disabled={data && currentPage * itemsPerPage >= data.length}
                        >
                            Next
                        </button>
                    </div>


                </div>


                {isLoading ? (
                    <p>Loading...</p>
                ) : <table className="table w-full">
                    <thead>
                    <tr>
                        <th>
                            <label>
                                <input type="checkbox" className="checkbox"/>
                            </label>
                        </th>
                        <th>Name</th>
                        <th>Job</th>
                        <th>Joined at</th>
                        <th></th>
                    </tr>
                    </thead>
                    <tbody>
                    {users?.map((user, index) => (
                        <tr key={index}>
                            <th>
                                <label>
                                    <input type="checkbox" className="checkbox bg-primary-100"/>
                                </label>
                            </th>
                            <td>
                                <div className="flex items-center gap-3">
                                    <div className="avatar text-gray-400">
                                        <div className="mask mask-squircle h-12 w-12">
                                            <img src={user?.avatar || "/boy.png"} alt="Avatar"/>
                                        </div>
                                    </div>
                                    <div>
                                        <div className="font-bold text-gray-200">{user?.username}</div>
                                        <div className="text-sm text-gray-400">{user.email}</div>
                                    </div>
                                </div>
                            </td>
                            <td>
                                {user.job}
                                <br/>
                                <span className="badge badge-ghost badge-sm">{user.jobTitle}</span>
                            </td>
                            <td>
                                <span className="  text-gray-400">
                                    {new Date(user.created_at).toLocaleString()}
                                </span>
                            </td>
                            <th>
                                <button className="btn btn-ghost btn-xs text-gray-400">details</button>
                            </th>
                        </tr>
                    ))}
                    </tbody>

                </table>}


            </div>
        </div>
    );
};

export default UsersList;
