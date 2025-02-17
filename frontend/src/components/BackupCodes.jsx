import React from 'react';

const BackupCodes = ({ recovery_codes = [] }) => {
    const handleDownload = () => {
        const blob = new Blob([recovery_codes.join('\n')], { type: 'text/plain' });
        const url = URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = 'backup_codes.txt';
        a.click();
        URL.revokeObjectURL(url);
    };

    console.log(recovery_codes)

    return (
        <div className="mt-6">
            {Array.isArray(recovery_codes) && (
                <>
                    <div className="flex items-center justify-between">
                        <h4 className="text-lg font-semibold text-gray-200">Backup Codes</h4>
                        <button
                            onClick={handleDownload}
                            className="btn btn-primary text-white px-4 py-2 rounded-lg"
                        >
                            Download
                        </button>
                    </div>
                    <ul className="mt-4 flex flex-wrap gap-2">
                        {recovery_codes?.map((code, index) => (
                            <li
                                key={index}
                                className="bg-gray-700 p-2 rounded-lg font-mono text-gray-100"
                            >
                                {code}
                            </li>
                        ))}
                    </ul>
                </>
            )}
        </div>
    );
};

export default BackupCodes;
