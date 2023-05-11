# Video-Upload-System


# Problem Statement ->
Design a video upload system for a user with low network bandwidth
User has to upload video which is more than 1GB. Users network bandwidth is too low. Network get dropped after 50% upload. User tries again and same thing happens. Now to design an optimized efficient solution to address this issue.


# Solution ->

To design an efficient solution for the given problem, we can implement a resumable video upload system with the help of chunked file uploading and an intelligent retry mechanism. Here is how it can be done:

Split the video into small chunks: We can split the video into small chunks of, say, 10MB or 20MB size. This will allow the user to upload the video in parts, and even if the upload fails, the user can resume the upload from the last successful chunk.

Use a retry mechanism: In case of network failure, we can implement an intelligent retry mechanism. The system can try uploading the failed chunks multiple times, with some delay between each attempt. If a chunk fails to upload even after multiple attempts, the system can mark that chunk as failed and move on to the next chunk.

Store metadata: The system can store metadata about the video, including the chunks that have been uploaded successfully, the chunks that have failed, and the total number of chunks. This metadata can be used to resume the upload from where it left off.

Checksum verification: After each chunk upload, the system can verify the checksum of the uploaded chunk to ensure data integrity. If the checksum doesn't match, the system can retry the upload of that chunk.

Notify the user: The system can notify the user about the progress of the upload, including the percentage of the video uploaded, the chunks uploaded successfully, and the failed chunks. This will keep the user informed and help them track the progress of the upload.

Provide an option to cancel the upload: In case the user wants to cancel the upload, we can provide an option to cancel the upload and delete all the uploaded chunks.

Optimize chunk size: We can optimize the chunk size based on the user's network speed to ensure optimal upload speed without overwhelming the user's network connection.

Implement a pause and resume feature: The system can provide a feature to pause and resume the upload process. The user can pause the upload and resume it at a later time.

By implementing the above steps, we can design an efficient video upload system for users with low network bandwidth.
